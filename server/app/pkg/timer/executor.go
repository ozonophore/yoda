package timer

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/configuration"
	jobf "github.com/yoda/app/pkg/job"
	"github.com/yoda/app/pkg/repository"
	service "github.com/yoda/app/pkg/service/logload"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"time"
)

type onBeforeJobExecution func(job *model.Job, transactionID int64, gJob *gocron.Job)
type onAfterJobExecution func(job *model.Job, transactionID int64, err error, gJob *gocron.Job)

func callOnBeforeJonExecution(job *model.Job, transactionID int64, gJob *gocron.Job, onBefore onBeforeJobExecution) {
	if onBefore != nil {
		onBefore(job, transactionID, gJob)
	}
}

func callOnAfterJonExecution(job *model.Job, transactionID int64, gJob *gocron.Job, err error, onAfter onAfterJobExecution) {
	if err == nil {
		logrus.Debugf("Job %d finished successfully", job.ID)
		repository.EndOperation(transactionID, types.StatusTypeCompleted)
	} else {
		logrus.Errorf("Job %d finished with error: %s", job.ID, err)
		repository.EndOperationWithMessage(transactionID, types.StatusTypeRejected, err.Error())
	}
	logrus.Infof("Finish job: %d. Next run: %s", job.ID, gJob.NextRun())
	if onAfter != nil {
		onAfter(job, transactionID, err, gJob)
	}
}

func jobByTag(s *gocron.Scheduler, jobId int) *gocron.Job {
	tag := fmt.Sprintf("%d", jobId)
	jobs, err := s.FindJobsByTag(tag)
	if err != nil {
		return nil
	}
	if len(jobs) == 0 {
		return nil
	}
	return jobs[0]
}

func RunRegularLoad(config *configuration.Config, ctx context.Context, jobID int, s *gocron.Scheduler, onBefore onBeforeJobExecution, onAfter onAfterJobExecution) {
	job, err := repository.GetJobWithOwnerByJobId(jobID)
	if err != nil {
		logrus.Errorf("Error after get jobs: %s with id: %d", err, jobID)
		return
	}
	if job.IsActive == false {
		logrus.Infof("Job %d is not active", jobID)
		return
	}
	transactionID := repository.BeginOperation(jobID)
	gJob := jobByTag(s, jobID)
	if gJob == nil {
		err = errors.New(fmt.Sprintf("Job %d not found", jobID))
		return
	}
	callOnBeforeJonExecution(job, transactionID, gJob, onBefore)
	err = execute(config, ctx, jobID, job, err, transactionID)
	callOnAfterJonExecution(job, transactionID, gJob, err, onAfter)
}

func execute(config *configuration.Config, ctx context.Context, jobID int, job *model.Job, err error, transactionID int64) error {
	logrus.Info("Start parsing for job: ", jobID)
	for _, param := range *job.Params {
		err = prepareParam(ctx, config, &param, transactionID)
		if err != nil {
			logrus.Errorf("Error after prepare param: %s", err)
			return err
		}
	}
	err = repository.CallDailyData(transactionID)
	if err != nil {
		logrus.Errorf("Error after call daily data: %s", err)
		return err
	}
	return nil
}

func prepareParam(ctx context.Context, config *configuration.Config, param *model.OwnerMarketplace, transactionID int64) error {
	service.CreateLogLoad(transactionID, param.OwnerCode, param.Source)
	var err error
	defer func(err error) {
		if err == nil {
			service.CompleteLogLoad(transactionID, param.OwnerCode, param.Source)
		} else {
			service.ErrorLogLoad(transactionID, param.OwnerCode, param.Source, err)
		}
	}(err)
	var loader jobf.DataLoader
	loader, err = jobf.JobFactory(param.Source, param.OwnerCode, *param.Password, param.ClientID, config)
	if err != nil {
		logrus.Panicf("Error after lookup a loader: %s", err)
	}
	newContext, _ := context.WithTimeout(ctx, time.Duration(config.Timeout)*time.Second)
	err = loader.Parsing(newContext, transactionID)
	return err
}

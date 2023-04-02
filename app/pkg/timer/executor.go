package timer

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/configuration"
	jobf "github.com/yoda/app/pkg/job"
	"github.com/yoda/app/pkg/repository"
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
		repository.EndOperation(transactionID, types.StatusTypeCompleted)
	} else {
		repository.EndOperation(transactionID, types.StatusTypeRejected)
	}
	logrus.Infof("Finish job: %d. Next run: %v", job.ID, gJob.NextRun())
	if onAfter != nil {
		onAfter(job, transactionID, err, gJob)
	}
}

func jobByTag(s *gocron.Scheduler, jobId int) *gocron.Job {
	tag := fmt.Sprintf("%d", jobId)
	jobs, err := s.FindJobsByTag(tag)
	if err != nil {
		logrus.Panicf("Error after find jobs by tag: %s", err)
		return nil
	}
	if len(jobs) == 0 {
		logrus.Panicf("No jobs found by tag: %d", jobId)
		return nil
	}
	return jobs[0]
}

func RunEtlJob(config *configuration.Config, ctx context.Context, jobID int, s *gocron.Scheduler, onBefore onBeforeJobExecution, onAfter onAfterJobExecution) {
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
	s.FindJobsByTag(fmt.Sprintf(`%s`, jobID))
	gJob := jobByTag(s, jobID)
	callOnBeforeJonExecution(job, transactionID, gJob, onBefore)
	defer callOnAfterJonExecution(job, transactionID, gJob, err, onAfter)
	logrus.Info("Start parsing for job: ", jobID)
	for _, param := range *job.Params {
		var loader jobf.DataLoader
		loader, err = jobf.JobFactory(param.Source, param.OwnerCode, *param.Password, param.ClientID, config)
		if err != nil {
			logrus.Errorf("Error after lookup a loader: %s", err)
			continue
		}
		newContext, _ := context.WithTimeout(ctx, time.Duration(config.Timeout)*time.Second)
		err = loader.Parsing(newContext, transactionID)
		if err != nil {
			logrus.Errorf("Error after parsing: %s", err)
			continue
		}
	}
}

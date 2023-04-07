package timer

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/model"
	"time"
)

type JobFunc = func(config *configuration.Config, ctx context.Context, jobID int)

func handleBeforeJobExecution(job *model.Job, transactionID int64, gJob *gocron.Job)           {}
func handleAfterJobExecution(job *model.Job, transactionID int64, err error, gJob *gocron.Job) {}

func InitScheduler(ctx *context.Context, config *configuration.Config) *gocron.Scheduler {
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.SingletonModeAll()

	jobs, err := repository.GetJobs()
	if err != nil {
		logrus.Panicf("Error after get jobs: %s", err)
	}
	for _, job := range *jobs {
		s := scheduler.Every(20)
		//if job.WeekDays == nil || job.WeekDays == nil {
		//	logrus.Panicf("Job %s has no week days", job.ID)
		//}
		//_, err := PrepareWeekDay(*job.WeekDays, s).At(job.AtTime).Do(RunEtlJob, config, *ctx, job.ID, handleBeforeJobExecution, handleAfterJobExecution)
		_, err := s.Second().Tag(fmt.Sprintf("%d", job.ID)).Do(RunEtlJob, config, *ctx, job.ID, scheduler, handleBeforeJobExecution, handleAfterJobExecution)
		if err != nil {
			logrus.Panicf("Error after create job: %s", err)
		}
	}
	scheduler.StartAsync()
	return scheduler
}

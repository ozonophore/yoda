package timer

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/repository"
	"github.com/yoda/common/pkg/dao"
	"github.com/yoda/common/pkg/model"
	"strings"
	"time"
)

type Scheduler struct {
	Scheduler *gocron.Scheduler
	config    *configuration.Config
}

type JobFunc = func(config *configuration.Config, ctx context.Context, jobID int)

func handleBeforeJobExecution(job *model.Job, transactionID int64, gJob *gocron.Job)           {}
func handleAfterJobExecution(job *model.Job, transactionID int64, err error, gJob *gocron.Job) {}

func NewScheduler(config *configuration.Config) *Scheduler {
	scheduler := gocron.NewScheduler(time.Local)
	dao.CreateScheduler(model.SCHEEDULER_MAIN)
	scheduler.SingletonModeAll()
	return &Scheduler{
		Scheduler: scheduler,
		config:    config,
	}
}

func (s *Scheduler) InitJob() {
	config := s.config
	ctx := context.Background()
	scheduler := s.Scheduler
	jobs, err := repository.GetJobs()
	if err != nil {
		logrus.Panicf("Error after get jobs: %s", err)
	}
	for _, job := range *jobs {
		switch strings.ToLower(job.Type) {
		case "interval":
			continue
		}
		if job.Type == "REGULAR" && job.WeekDays == nil {
			logrus.Panicf("Job %s has no week days", job.ID)
		}
		s := s.Scheduler.Every(1)
		_, err := PrepareWeekDay(*job.WeekDays, s).At(job.AtTime).Do(RunRegularLoad, config, ctx, job.ID, scheduler, handleBeforeJobExecution, handleAfterJobExecution)
		//_, err := s.Second().Tag(fmt.Sprintf("%d", job.ID)).Do(RunRegularLoad, config, ctx, job.ID, scheduler, handleBeforeJobExecution, handleAfterJobExecution)
		if err != nil {
			logrus.Panicf("Error after create job: %s", err)
		}
	}
}

func (s *Scheduler) Start() {
	s.Scheduler.StartAsync()
	dao.UpdateScheduler(model.SCHEEDULER_MAIN, model.STATUS_RUNNING)
	logrus.Info("Scheduler started")
}

func (s *Scheduler) Stop() {
	s.Scheduler.Stop()
	dao.UpdateScheduler(model.SCHEEDULER_MAIN, model.STATUS_STOPPED)
	logrus.Info("Scheduler stopped")
}

func InitScheduler(ctx *context.Context, config *configuration.Config) *gocron.Scheduler {
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.SingletonModeAll()

	jobs, err := repository.GetJobs()
	if err != nil {
		logrus.Panicf("Error after get jobs: %s", err)
	}
	for _, job := range *jobs {
		switch strings.ToLower(job.Type) {
		case "interval":
			continue
		}
		s := scheduler.Every(20)
		//if job.WeekDays == nil || job.WeekDays == nil {
		//	logrus.Panicf("Job %s has no week days", job.ID)
		//}
		//_, err := PrepareWeekDay(*job.WeekDays, s).At(job.AtTime).Do(RunRegularLoad, config, *ctx, job.ID, handleBeforeJobExecution, handleAfterJobExecution)
		_, err := s.Second().Tag(fmt.Sprintf("%d", job.ID)).Do(RunRegularLoad, config, *ctx, job.ID, scheduler, handleBeforeJobExecution, handleAfterJobExecution)
		if err != nil {
			logrus.Panicf("Error after create job: %s", err)
		}
	}
	scheduler.StartAsync()
	return scheduler
}

package timer

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/configuration"
	"github.com/yoda/app/pkg/event"
	"github.com/yoda/common/pkg/dao"
	"github.com/yoda/common/pkg/model"
	"time"
)

type Scheduler struct {
	scheduler       *gocron.Scheduler
	systemScheduler *gocron.Scheduler
	config          *configuration.Config
}

type JobFunc = func(config *configuration.Config, ctx context.Context, jobID int)

func handleBeforeJobExecution(job *model.Job, transactionID int64, gJob *gocron.Job)           {}
func handleAfterJobExecution(job *model.Job, transactionID int64, err error, gJob *gocron.Job) {}

func NewScheduler(config *configuration.Config) *Scheduler {
	scheduler := gocron.NewScheduler(time.Local)
	dao.CreateScheduler(model.SCHEEDULER_MAIN)
	scheduler.SingletonModeAll()

	system := gocron.NewScheduler(time.Local)
	system.SingletonModeAll()
	return &Scheduler{
		scheduler:       scheduler,
		systemScheduler: system,
		config:          config,
	}
}

func (s *Scheduler) GetAllJobs() []*gocron.Job {
	return append(s.scheduler.Jobs(), s.systemScheduler.Jobs()...)
}

func (s *Scheduler) RunImmediately(jobID int) {
	tag := fmt.Sprintf(`%d`, jobID)
	err := s.scheduler.RunByTag(tag)
	if err != nil {
		logrus.Errorf("Error after run job: %s", err)
	}
}

func (s *Scheduler) InitJob() {
	ctx := context.Background()
	if s.initRegular(ctx) {
		event.RefreshJobs(GetJobsFromCache())
	}
	s.initSystem(ctx)
}

func (s *Scheduler) Start() {
	s.scheduler.StartAsync()
	_, t := s.scheduler.NextRun()
	logrus.Debugf("Next run: %v", t)
	dao.UpdateScheduler(model.SCHEEDULER_MAIN, model.STATUS_RUNNING)
	logrus.Info("scheduler started")
}

func (s *Scheduler) Stop() {
	s.scheduler.Stop()
	dao.UpdateScheduler(model.SCHEEDULER_MAIN, model.STATUS_STOPPED)
	logrus.Info("scheduler stopped")
}

func (s *Scheduler) StopAll() {
	s.Stop()
	s.systemScheduler.Stop()
}

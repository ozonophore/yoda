package timer

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/configuration"
	"github.com/yoda/app/internal/event"
)

var (
	logger *logrus.Logger
)

func (s Scheduler) initSystem(ctx context.Context, config *configuration.System) {
	if logger == nil {
		logger = logrus.New()
		loglevel, _ := logrus.ParseLevel(config.LoggingLevel)
		logger.SetLevel(loglevel)
	}
	s.systemScheduler.Every(1).Minute().Do(func() {
		logger.Debug("Refresh jobs")
		wasChanged := s.initRegular(ctx)
		if wasChanged {
			logger.Debug("Jobs was changed")
			event.RefreshJobs(GetJobsFromCache())
		}
	})
	s.systemScheduler.StartAsync()
	//scheduler := gocron.NewScheduler(time.UTC)
	//scheduler.Every(1).Day().At("00:00").Do(DoSystem)
	//scheduler.StartAsync()
}

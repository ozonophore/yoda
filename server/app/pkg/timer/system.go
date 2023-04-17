package timer

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/pkg/event"
)

func (s *Scheduler) initSystem(ctx context.Context) {
	s.systemScheduler.Every(1).Minute().Do(func() {
		logrus.Debug("Refresh jobs")
		wasChanged := s.initRegular(ctx)
		if wasChanged {
			logrus.Debug("Jobs was changed")
			event.RefreshJobs(GetJobsFromCache())
		}
	})
	s.systemScheduler.StartAsync()
	//scheduler := gocron.NewScheduler(time.UTC)
	//scheduler.Every(1).Day().At("00:00").Do(DoSystem)
	//scheduler.StartAsync()
}

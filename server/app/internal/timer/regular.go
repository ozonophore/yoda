package timer

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/repository"
	"strings"
)

func (sch Scheduler) initRegular(ctx context.Context) bool {
	wasChanged := false
	logrus.Debugf("Init jobs")
	config := sch.config
	scheduler := sch.scheduler

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
			logrus.Panicf("Job %d has no week days", job.ID)
		}
		tag := fmt.Sprintf(`%d`, job.ID)

		ok := ExistsInCache(job.ID)
		if ok {
			isEquals := Equals(&job)
			if isEquals {
				continue
			}
			sch.scheduler.RemoveByTag(tag)
		}
		AddJobToCache(&job, scheduler)

		s := sch.scheduler.Every(1)
		atTime := strings.ReplaceAll(*job.AtTime, ",", ";")
		logrus.Debugf("Init job: %d, %s, %s", job.ID, *job.WeekDays, atTime)
		job, err := PrepareWeekDay(*job.WeekDays, s).At(atTime).Tag(tag).Do(RunRegularLoad, config, ctx, job.ID, scheduler, sch.handleBeforeJobExecution, sch.handleAfterJobExecution)
		logrus.Debugf("Next run: %s", job.NextRun())
		if err != nil {
			logrus.Panicf("Error after create job: %s", err)
		}
		wasChanged = true
	}
	return wasChanged
}

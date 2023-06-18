package stage

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/pipeline"
	time2 "github.com/yoda/app/internal/time"
	"github.com/yoda/common/pkg/model"
	"sort"
	"strings"
	"time"
)

type stageFactory interface {
	CreateStage(jobId int) (pipeline.Stage, error)
}

type stageRepository interface {
	GetJob(id int) (*model.Job, error)
}

type Initializer struct {
	rep          stageRepository
	stageFactory stageFactory
	sch          *gocron.Scheduler
}

func NewInitializer(rep stageRepository, stageFactory stageFactory, sch *gocron.Scheduler) *Initializer {
	return &Initializer{rep: rep, stageFactory: stageFactory, sch: sch}
}

func (s *Initializer) jobByTag(jobId int) *gocron.Job {
	tag := fmt.Sprintf("%d", jobId)
	jobs, err := s.sch.FindJobsByTag(tag)
	if err != nil {
		return nil
	}
	if len(jobs) == 0 {
		return nil
	}
	return jobs[0]
}

func (s *Initializer) check(jobId int) (bool, *string, *string, error) {
	const ok = false
	j, err := s.rep.GetJob(jobId)
	if !j.IsActive {
		return ok, nil, nil, fmt.Errorf("job with tag(%d) is not active", jobId)
	}
	if err != nil {
		logrus.Panicf("Error while getting job(%d): %v", 2, err)
	}
	job := s.jobByTag(jobId)
	atTime := strings.ReplaceAll(*j.AtTime, ",", ";")
	var atWeekDays = ""
	if j.WeekDays != nil {
		atWeekDays = strings.ReplaceAll(*j.WeekDays, ",", ";")
	}

	if job != nil {
		logrus.Infof("Job with tag(%d) already exists", jobId)
		weekDaysEquals := false
		if len(atWeekDays) > 0 {
			actualWeekdays := time2.WeekdayToArray(job.Weekdays())
			expectedWeekDays := strings.Split(atWeekDays, ";")
			sort.Strings(actualWeekdays)
			sort.Strings(expectedWeekDays)
			weekDaysEquals = strings.Join(actualWeekdays, ";") == strings.Join(expectedWeekDays, ";")
		} else {
			weekDaysEquals = true
		}

		actualAtTimes := job.ScheduledAtTimes()
		expectedAtTimes := strings.Split(atTime, ";")

		sort.Strings(actualAtTimes)
		sort.Strings(expectedAtTimes)
		atTimeEquals := strings.Join(actualAtTimes, ";") == strings.Join(expectedAtTimes, ";")
		if atTimeEquals && weekDaysEquals {
			return ok, nil, nil, nil
		}
		err := s.sch.RemoveByTag(fmt.Sprintf("%d", jobId))
		if err != nil {
			return ok, nil, nil, fmt.Errorf("error while removing job with tag(%d): %v", jobId, err)
		}
	}
	return !ok, &atWeekDays, &atTime, nil
}

func (s *Initializer) createJobAtTime(tag, atTime string, stage pipeline.Stage) *gocron.Job {
	job, err := s.sch.Every(1).Day().At(atTime).Tag(tag).Do(func() {
		p := pipeline.NewPipeline()
		err := p.Do(context.Background(), stage).Error()
		if err != nil {
			logrus.Errorf("Error while running stock aggregator: %v", err)
		}
	})
	if err != nil {
		logrus.Panicf("Error while scheduling job: %v", err)
	}
	return job
}

func (s *Initializer) createJobWeekdaysAndAtTime(tag, atTime string, weekDays []time.Weekday, stage pipeline.Stage) *gocron.Job {
	sch := s.sch.Every(1)
	for i := 0; i < len(weekDays); i++ {
		sch.Weekday(weekDays[i])
	}
	job, err := s.sch.Every(1).At(atTime).Tag(tag).Do(func() {
		p := pipeline.NewPipeline()
		err := p.Do(context.Background(), stage)
		if err != nil {
			logrus.Errorf("Error while running stock aggregator: %v", err)
		}
	})
	if err != nil {
		logrus.Panicf("Error while scheduling job: %v", err)
	}
	return job
}

func (s *Initializer) Do(jobId int) error {
	ok, weekdays, atTime, err := s.check(jobId)
	if err != nil {
		return fmt.Errorf("error while checking job(%d): %v", jobId, err)
	}
	if !ok {
		return nil
	}
	logrus.Debugf("check was passed for job(%d) weekdays(%s) and atTime(%s)", jobId, *weekdays, *atTime)
	stage, err := s.stageFactory.CreateStage(jobId)
	if err != nil {
		return fmt.Errorf("error while creating stage for job(%d): %v", jobId, err)
	}
	tag := fmt.Sprintf("%d", jobId)
	var job *gocron.Job
	if (weekdays != nil) && (len(*weekdays) > 0) {
		days := strings.Split(*weekdays, ";")
		wds := time2.StringsToWeekdays(days)
		job = s.createJobWeekdaysAndAtTime(tag, *atTime, wds, stage)
	} else {
		job = s.createJobAtTime(tag, *atTime, stage)
	}
	job.SetEventListeners(func() {
		logrus.Infof("Job with tag(%d) is running", 2)
	}, func() {
		logrus.Infof("Job with tag(%d) is done. Next run: %s", 2, job.NextRun().UTC())
	})
	return nil
}

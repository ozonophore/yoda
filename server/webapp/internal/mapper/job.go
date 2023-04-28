package mapper

import (
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/webapp/internal/api"
	"strings"
)

func MapJobsToApi(s []model.Job) *api.Job {
	var jobs = make(map[int]*model.Job)
	for _, sjob := range s {
		jobs[sjob.ID] = &sjob
	}
	job := api.Job{
		Id:             0,
		Loader:         MapJobToLoader(jobs[1]),
		CalcAggregates: MapJobToLoader(jobs[2]),
		AddLoader:      MapJobToAddLoader(jobs[3]),
	}
	return &job
}

func MapJobToLoader(job *model.Job) api.JobLoader {
	if job == nil {
		return api.JobLoader{
			WeekDays: []api.WeekDay{},
			AtTimes:  []string{},
			NextRun:  nil,
			LastRun:  nil,
		}
	}
	return api.JobLoader{
		WeekDays: MapArrayToWeekDays(job.WeekDays),
		AtTimes:  strings.Split(*job.AtTime, ","),
		NextRun:  job.NextRun,
		LastRun:  job.LastRun,
	}
}

func MapJobToAddLoader(job *model.Job) api.JobAddLoader {
	if job == nil {
		return api.JobAddLoader{
			Interval: nil,
			MaxRuns:  nil,
		}
	}
	return api.JobAddLoader{
		Interval: job.Interval,
		MaxRuns:  job.MaxRuns,
	}
}

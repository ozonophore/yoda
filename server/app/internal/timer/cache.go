package timer

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/yoda/common/pkg/eventbus"
	"github.com/yoda/common/pkg/model"
	"sync"
)

type CachedJob struct {
	ID          string
	Description string
	Scheduler   *gocron.Scheduler
	WeekDays    *string
	AtTime      *string
}

var cacheJobs = make(map[int]*CachedJob)
var mutex sync.Mutex

func ExistsInCache(jobID int) bool {
	_, ok := cacheJobs[jobID]
	return ok
}

func AddJobToCache(job *model.Job, scheduler *gocron.Scheduler) {
	mutex.Lock()
	defer mutex.Unlock()
	cacheJobs[job.ID] = &CachedJob{
		ID:          fmt.Sprintf(`%d`, job.ID),
		Description: *job.Description,
		Scheduler:   scheduler,
		WeekDays:    job.WeekDays,
		AtTime:      job.AtTime,
	}
}

func Equals(job *model.Job) bool {
	mutex.Lock()
	defer mutex.Unlock()
	cachedJob := cacheJobs[job.ID]
	return equalsString(job.Description, &cachedJob.Description) &&
		equalsString(job.WeekDays, cachedJob.WeekDays) &&
		equalsString(job.AtTime, cachedJob.AtTime)
}

func equalsString(a *string, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func GetJobsFromCache() []eventbus.MQJob {
	mutex.Lock()
	defer mutex.Unlock()
	var jobs = make([]eventbus.MQJob, len(cacheJobs))
	index := 0
	for _, value := range cacheJobs {
		jobs[index] = mapCachedJobToJob(*value)
		index++
	}
	return jobs
}

func mapCachedJobToJob(cachedJob CachedJob) eventbus.MQJob {
	job, _ := cachedJob.Scheduler.FindJobsByTag(cachedJob.ID)
	return eventbus.MQJob{
		ID:          cachedJob.ID,
		Description: cachedJob.Description,
		WeekDays:    ParseStringToArray(cachedJob.WeekDays),
		AtTimes:     ParseStringToArray(cachedJob.AtTime),
		IsRunning:   len(job) > 0 && job[0].IsRunning(),
	}
}

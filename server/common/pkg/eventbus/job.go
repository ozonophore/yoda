package eventbus

import "time"

type MQJob struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Interval    *int      `json:"interval"`
	runCount    *int      `json:"run_count"`
	WeekDays    []string  `json:"weekdays"`
	AtTimes     []string  `json:"attimes"`
	LastRun     time.Time `json:"lastrun"`
	NextRun     time.Time `json:"nextrun"`
	IsRunning   bool      `json:"is_running"`
}

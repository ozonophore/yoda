package event

import (
	"encoding/json"
	"github.com/yoda/common/pkg/eventbus"
)

func AddJob(id, description string, weekday, atTime []string, isRunning bool) {
	event := eventbus.MQJob{
		ID:          id,
		Description: description,
		WeekDays:    weekday,
		AtTimes:     atTime,
		IsRunning:   isRunning,
	}
	body, _ := json.Marshal(event)
	eventbus.Publish("scheduler", eventbus.EVENT_JOB_ADD, "0", body)
}

package timer

import (
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"strings"
)

func ParseStringToArray(s *string) []string {
	if s == nil {
		return []string{}
	}
	return strings.Split(strings.ReplaceAll(*s, ";", ","), ",")
}

func PrepareWeekDay(s string, scheduler *gocron.Scheduler) *gocron.Scheduler {
	items := ParseStringToArray(&s)
	for _, day := range items {
		switch strings.ToLower(strings.TrimSpace(day)) {
		case "monday":
			scheduler.Monday()
		case "tuesday":
			scheduler.Tuesday()
		case "wednesday":
			scheduler.Wednesday()
		case "thursday":
			scheduler.Thursday()
		case "friday":
			scheduler.Friday()
		case "saturday":
			scheduler.Saturday()
		case "sunday":
			scheduler.Sunday()
		default:
			logrus.Error("invalid day")
		}
	}
	return scheduler
}

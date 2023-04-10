package mapper

import (
	"github.com/yoda/webapp/pkg/api"
	"strings"
)

func MapArrayToWeekDays(s *string) []api.WeekDay {
	days := strings.Split(*s, ",")
	count := 0
	if s != nil {
		count = len(days)
	}
	newDays := make([]api.WeekDay, count)
	for i, day := range days {
		newDays[i] = api.WeekDay(day)
	}
	return newDays
}

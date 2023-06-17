package time

import (
	"k8s.io/utils/strings/slices"
	"time"
)

var longDayNames = []string{
	"sunday",
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
}

func WeekdayToArray(days []time.Weekday) []string {
	l := len(days)
	if l == 0 {
		return []string{}
	}
	s := make([]string, l, l)
	for i := 0; i < l; i++ {
		s[i] = longDayNames[days[i]]
	}
	return s
}

func StringsToWeekdays(days []string) []time.Weekday {
	l := len(days)
	wd := make([]time.Weekday, l, l)
	for i := 0; i < l; i++ {
		wd[i] = time.Weekday(slices.Index(longDayNames, days[i]))
	}
	return wd
}

package time

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWeekdayToArray(t *testing.T) {
	days := []time.Weekday{time.Monday, time.Wednesday}
	target := WeekdayToArray(days)
	assert.Equal(t, 2, len(target))
	assert.Equal(t, "monday", target[0])
	assert.Equal(t, "wednesday", target[1])
}

func TestStringsToWeekdays(t *testing.T) {
	days := []string{"tuesday", "thursday"}
	target := StringsToWeekdays(days)
	assert.Equal(t, 2, len(target))
	assert.Equal(t, time.Weekday(2), target[0])
	assert.Equal(t, time.Weekday(4), target[1])
}

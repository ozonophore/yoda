package types

import (
	"fmt"
	"strings"
	"time"
)

const ctLayout = time.DateOnly + "T" + time.TimeOnly

type CustomTime time.Time

// UnmarshalJSON Parses the json string in the custom format
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.ReplaceAll(strings.Trim(string(b), `"`), "Z", "")
	nt, err := time.Parse(ctLayout, s)
	if err != nil {
		return err
	}
	*ct = CustomTime(nt)
	return
}

// ToTime returns the time.Time value
func (ct *CustomTime) ToTime() time.Time {
	return time.Time(*ct)
}

// CustomTimeToTime returns the time.Time value
func CustomTimeToTime(ct *CustomTime) *time.Time {
	if ct == nil {
		return nil
	}
	v := time.Time(*ct)
	return &v
}

// MarshalJSON writes a quoted string in the custom format
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

// String returns the time in the custom format
func (ct *CustomTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf(`%q`, t.Format(ctLayout))
}

func (ct *CustomTime) ToString() string {
	t := time.Time(*ct)
	return t.Format(ctLayout)
}

func StringToCustomTime(s string) *CustomTime {
	if s == "" {
		return nil
	}
	t, err := time.Parse(ctLayout, s)
	if err != nil {
		return nil
	}
	ct := CustomTime(t)
	return &ct
}

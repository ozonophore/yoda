package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
)

func Float32ToFloat64(s *float32) *float64 {
	if s == nil {
		return nil
	}
	t := float64(*s)
	return &t
}

func IntToInt32(s *int) *int32 {
	if s == nil {
		return nil
	}
	t := int32(*s)
	return &t
}

func Int64ToInt32(s *int64) *int32 {
	if s == nil {
		return nil
	}
	t := int32(*s)
	return &t
}

func IntToBoolean(s *int) bool {
	if s == nil {
		return false
	}
	if *s == 1 {
		return true
	}
	return false
}

func Float64ToFloat32(s *float64) *float32 {
	if s == nil {
		return nil
	}
	t := float32(*s)
	return &t
}

func BooleanToBoolean(s *bool) bool {
	if s == nil {
		return false
	}
	return *s
}

func StringToFloat64(value *string) float64 {
	if value == nil {
		return 0
	}
	result, err := strconv.ParseFloat(*value, 64)
	if err != nil {
		str, _ := json.Marshal(*value)
		logrus.Errorf("Error parse price %s", str)
		return 0
	}
	return result
}

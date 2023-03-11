package utils

func Float32ToFloat64(s *float32) *float64 {
	t := float64(*s)
	return &t
}

func IntToInt32(s *int) *int32 {
	t := int32(*s)
	return &t
}

func Int64ToInt32(s *int64) *int32 {
	t := int32(*s)
	return &t
}

func IntToBoolean(s *int) bool {
	if *s == 1 {
		return true
	}
	return false
}

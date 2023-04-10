package utils

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

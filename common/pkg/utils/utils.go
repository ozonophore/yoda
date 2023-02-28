package utils

func Float32ToFloat64(s *float32) *float64 {
	t := float64(*s)
	return &t
}

func IntToInt32(s *int) *int32 {
	t := int32(*s)
	return &t
}

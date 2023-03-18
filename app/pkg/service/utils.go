package service

import "math"

type CallbackFunc[T any] func(items *[]T) error

func CallbackBatch[T any](items *[]T, batchSize int, callback CallbackFunc[T]) error {
	var low int
	low = 0
	highest := len(*items)
	for {
		step := int(math.Min(float64(batchSize), float64(highest-low)))
		high := low + step
		batches := (*items)[low:high]
		err := callback(&batches)
		if err != nil {
			return err
		}
		low = high
		if high == highest {
			break
		}
	}
	return nil
}

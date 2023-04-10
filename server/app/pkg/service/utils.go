package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"gopkg.in/errgo.v2/fmt/errors"
	"math"
)

type CallbackFunc[T any] func(items *[]T) error

type FetchFunc func(offset int64, limit int64) (int64, error)

func CallbackBatch[T any](items *[]T, batchSize int, callback CallbackFunc[T]) error {
	var low int
	low = 0
	if items == nil {
		return errors.New("Null pointer exception: items is nil")
	}
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

func FetchBatch(context context.Context, batchSize int64, fetchFunc FetchFunc) error {
	var offset int64
	offset = 0
	for {
		select {
		case <-context.Done():
			logrus.Error("Process was canceled")
			break
		default:
			count, err := fetchFunc(offset, batchSize)
			if err != nil {
				return err
			}
			if count < batchSize {
				return nil
			}
			offset += batchSize
		}
	}
}

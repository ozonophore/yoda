package stock

import (
	"context"
	"github.com/yoda/app/internal/pipeline"
	"time"
)

type dailyService interface {
	CalcAggrByDay(day time.Time) error
}

type DailyStep struct {
	service dailyService
}

func NewDailyStep(service dailyService) *DailyStep {
	return &DailyStep{
		service: service,
	}
}

func (d *DailyStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	day := time.Now().UTC().Truncate(24 * time.Hour)
	return day, d.service.CalcAggrByDay(day)
}

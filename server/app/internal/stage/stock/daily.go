package stock

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/pipeline"
	"time"
)

type dailyService interface {
	CalcAggrByDay(day time.Time) error
}

type DailyStep struct {
	service dailyService
	logger  *logrus.Logger
}

func NewDailyStep(service dailyService, logg *logrus.Logger) *DailyStep {
	return &DailyStep{
		service: service,
		logger:  logg,
	}
}

func (d *DailyStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	day := time.Now().UTC().Truncate(24 * time.Hour)
	d.logger.Debugf("DailyStep: %s", day)
	return day, d.service.CalcAggrByDay(day)
}

package stock

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/pipeline"
	"time"
)

type defectureService interface {
	CalcDef(day time.Time) error
}

type DefectureStep struct {
	service defectureService
	logger  *logrus.Logger
}

func NewDefectureStep(service defectureService, logg *logrus.Logger) *DefectureStep {
	return &DefectureStep{
		service: service,
		logger:  logg,
	}
}

func (d *DefectureStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	ps, ok := (*deps)["stock-daily-aggregator"]
	if !ok {
		return nil, errors.New("stock-daily-aggregator not found")
	}
	status := ps.GetStatus()
	date := status.Value.(time.Time)
	return nil, d.service.CalcDef(date)
}

package stock

import (
	"context"
	"errors"
	"github.com/yoda/app/internal/pipeline"
	"time"
)

type defectureService interface {
	CalcDefFor30(day time.Time) error
}

type DefectureStep struct {
	service defectureService
}

func NewDefectureStep(service defectureService) *DefectureStep {
	return &DefectureStep{
		service: service,
	}
}

func (d *DefectureStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	ps, ok := (*deps)["stock-daily-aggregator"]
	if !ok {
		return nil, errors.New("stock-daily-aggregator not found")
	}
	status := ps.GetStatus()
	date := status.Value.(time.Time)
	return nil, d.service.CalcDefFor30(date)
}

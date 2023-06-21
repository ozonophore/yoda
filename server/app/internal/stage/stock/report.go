package stock

import (
	"context"
	"errors"
	"github.com/yoda/app/internal/pipeline"
	"time"
)

type reportService interface {
	CalcReport(day time.Time) error
}

type ReportStep struct {
	service reportService
}

func NewReportStep(service reportService) *ReportStep {
	return &ReportStep{
		service: service,
	}
}

func (d *ReportStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	ps, ok := (*deps)["stock-daily-aggregator"]
	if !ok {
		return nil, errors.New("stock-daily-aggregator not found")
	}
	status := ps.GetStatus()
	date := status.Value.(time.Time)
	return nil, d.service.CalcReport(date)
}

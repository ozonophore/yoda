package stock

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/pipeline"
	"time"
)

type reportService interface {
	CalcReport(day time.Time) error
}

type reportByCluesterService interface {
	CalcReportByClusters(day time.Time) error
}

type ReportStep struct {
	service reportService
	logger  *logrus.Logger
}

func NewReportStep(service reportService, logg *logrus.Logger) *ReportStep {
	return &ReportStep{
		service: service,
		logger:  logg,
	}
}

func (d *ReportStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	ps, ok := (*deps)["stock-daily-aggregator"]
	if !ok {
		return nil, errors.New("stock-daily-aggregator not found")
	}
	status := ps.GetStatus()
	date := status.Value.(time.Time)
	d.logger.Debugf("ReportStep: %s", date)
	return nil, d.service.CalcReport(date)
}

type ReportByClustersStep struct {
	service reportByCluesterService
	logger  *logrus.Logger
}

func NewReportByClustersStep(service reportByCluesterService, logg *logrus.Logger) *ReportByClustersStep {
	return &ReportByClustersStep{
		service: service,
		logger:  logg,
	}
}

func (d *ReportByClustersStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	ps, ok := (*deps)["stock-daily-aggregator"]
	if !ok {
		return nil, errors.New("stock-daily-aggregator not found")
	}
	status := ps.GetStatus()
	date := status.Value.(time.Time)
	d.logger.Debugf("ReporClustertStep: %s", date)
	return nil, d.service.CalcReportByClusters(date)
}

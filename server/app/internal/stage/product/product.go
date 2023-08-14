package product

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/pipeline"
	"time"
)

type productService interface {
	CalcDefAndReportByProduct(day time.Time) error
}

type ProductStep struct {
	service productService
	logger  *logrus.Logger
}

func NewProductStep(service productService, logg *logrus.Logger) *ProductStep {
	return &ProductStep{
		service: service,
		logger:  logg,
	}
}

func (d *ProductStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	ps, ok := (*deps)["stock-daily-aggregator"]
	if !ok {
		return nil, errors.New("stock-daily-aggregator not found")
	}
	status := ps.GetStatus()
	date := status.Value.(time.Time)
	d.logger.Debugf("Calcucaliton of report by the products was started")
	return nil, d.service.CalcDefAndReportByProduct(date)
}

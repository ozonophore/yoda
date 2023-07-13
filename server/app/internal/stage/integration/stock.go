package integration

import (
	"context"
	"github.com/sirupsen/logrus"
	integration "github.com/yoda/app/internal/integration/api"
	"github.com/yoda/app/internal/pipeline"
)

type IntegrationStockService interface {
	UploadStocks(ctx context.Context, stocks *[]integration.Stock) error
}

type StockStep struct {
	service IntegrationStockService
	logger  *logrus.Logger
	client  integration.ClientWithResponsesInterface
}

func NewStockStep(service IntegrationStockService, client integration.ClientWithResponsesInterface, logg *logrus.Logger) *StockStep {
	return &StockStep{
		service: service,
		logger:  logg,
		client:  client,
	}
}

func (d *StockStep) Do(ctx context.Context, deps *map[string]pipeline.Stage, e error) (interface{}, error) {
	resp, err := d.client.GetStocksWithResponse(ctx)
	if err != nil {
		d.logger.Errorf("Error while getting stocks: %v", err)
		return nil, err
	}
	if resp.JSON200 == nil {
		d.logger.Debug("No stocks to upload")
		return nil, nil
	}
	stocks := resp.JSON200
	d.logger.Debugf("Got %d stocks", len(stocks.Items))
	err = d.service.UploadStocks(ctx, &stocks.Items)
	if err != nil {
		d.logger.Errorf("Error while uploading stocks: %v", err)
	}
	return nil, err
}

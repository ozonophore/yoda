package integration

import (
	"github.com/sirupsen/logrus"
	integration "github.com/yoda/app/internal/integration/api"
	"github.com/yoda/app/internal/pipeline"
)

func NewStockStage(service IntegrationStockService, client integration.ClientWithResponsesInterface, logger *logrus.Logger) pipeline.Stage {
	step := NewStockStep(service, client, logger)
	stage := pipeline.NewSimpleStageWithTag(step, "int-stock")
	return stage
}

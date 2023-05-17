package dictionary

import (
	"context"
	"errors"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/sirupsen/logrus"
	integration "github.com/yoda/app/internal/integration/api"
	"github.com/yoda/app/internal/logging"
	"github.com/yoda/app/internal/storage"
	"github.com/yoda/common/pkg/model"
	"net/http"
	"time"
)

func UpdateMarketplaces() error {
	logrus.Debug("Start update Marketplaces")
	defer logrus.Debug("End update Marketplaces")
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Key", c.Token)
	client, err := integration.NewClientWithResponses(c.Host, logging.WithLoggerIntegrationFn(c.LogLevel), integration.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return errors.Join(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	resp, err := client.GetMarketplacesWithResponse(ctx)
	if err != nil {
		return errors.Join(err)
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New("invalid status code")
	}
	result := resp.JSON200
	if result == nil {
		return errors.New("invalid response")
	}
	if result.Count == 0 {
		return nil
	}
	items := result.Items
	var marketplaces = make([]model.Marketplace, len(items))
	for i, item := range items {
		marketplaces[i].ID = item.Id
		marketplaces[i].Name = item.Name
		marketplaces[i].UpdatedAt = item.UpdateAt.ToTime()
	}
	return errors.Join(storage.SaveOrUpdateMarketplaces(&marketplaces))
}

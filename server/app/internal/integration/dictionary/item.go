package dictionary

import (
	"context"
	"errors"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/configuration"
	integration "github.com/yoda/app/internal/integration/api"
	"github.com/yoda/app/internal/logging"
	"github.com/yoda/app/internal/storage"
	"github.com/yoda/common/pkg/model"
	"net/http"
	"time"
)

var c configuration.Integration

func InitDictionary(config configuration.Integration) {
	c = config
}

func UpdateItems() error {
	logrus.Debug("Start update items")
	defer logrus.Debug("End update items")
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Key", c.Token)
	client, err := integration.NewClientWithResponses(c.Host, logging.WithLoggerIntegrationFn(c.LogLevel), integration.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return errors.Join(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	resp, err := client.GetItemsWithResponse(ctx)
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
		logrus.Error("Items count is 0")
		return nil
	}
	actualCount := storage.GetItemCount()
	if int32(actualCount) >= result.Count {
		return nil
	}
	items := result.Items
	var newItems = make([]model.Item, len(items))
	for i, item := range items {
		newItems[i].ID = item.Id
		newItems[i].Name = item.Name
		newItems[i].UpdatedAt = item.UpdateAt.ToTime()
	}
	return errors.Join(storage.SaveOrUpdateItem(&newItems))
}

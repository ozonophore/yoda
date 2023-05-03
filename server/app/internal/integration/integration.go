package integration

import (
	"context"
	"errors"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/sirupsen/logrus"
	"github.com/yoda/app/internal/configuration"
	integration "github.com/yoda/app/internal/integration/api"
	"github.com/yoda/app/internal/logging"
	"github.com/yoda/app/internal/repository"
	"github.com/yoda/common/pkg/model"
	"net/http"
	"time"
)

var instance *UpdaterOrganisations

type UpdaterOrganisations struct {
	config configuration.Integration
}

func NewUpdaterOrganisations(config configuration.Integration) *UpdaterOrganisations {
	if instance == nil {
		instance = &UpdaterOrganisations{
			config: config,
		}
	}
	return instance
}

func InstanceUpdaterOrganisations() *UpdaterOrganisations {
	return instance
}

func (i *UpdaterOrganisations) UpdateOrganizations() error {
	logrus.Debug("Start update organizations")
	defer logrus.Debug("End update organizations")
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Key", i.config.Token)
	client, err := integration.NewClientWithResponses(i.config.Host, logging.WithLoggerIntegrationFn(i.config.LogLevel), integration.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return errors.Join(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(i.config.Timeout)*time.Second)
	resp, err := client.GetOrganizationsWithResponse(ctx)
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
	var organizations = make([]model.Organisation, len(items))
	for i, item := range items {
		organizations[i].ID = item.Id
		organizations[i].Name = item.Name
		organizations[i].Inn = item.Inn
		organizations[i].Kpp = item.Kpp
		organizations[i].UpdateAt = item.UpdateAt.ToTime()
	}
	return errors.Join(repository.SaveOrganisations(&organizations))
}

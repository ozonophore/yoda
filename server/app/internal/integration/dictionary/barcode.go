package dictionary

import (
	"context"
	"errors"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/sirupsen/logrus"
	integration "github.com/yoda/app/internal/integration/api"
	"github.com/yoda/app/internal/logging"
	"github.com/yoda/app/internal/repository"
	"github.com/yoda/common/pkg/model"
	"net/http"
	"time"
)

func UpdateBarcode() error {
	logrus.Debug("Start update barcodes")
	defer logrus.Debug("End update barcodes")
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Key", c.Token)
	client, err := integration.NewClientWithResponses(c.Host, logging.WithLoggerIntegrationFn(c.LogLevel), integration.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return errors.Join(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	resp, err := client.GetItemsBarcodesWithResponse(ctx)
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
	var barcodes = make([]model.Barcode, len(items))
	for i, item := range items {
		barcodes[i].ItemID = item.Id
		barcodes[i].BarcodeID = item.BarcodeID
		barcodes[i].Barcode = item.Barcode
		barcodes[i].OrganisationID = item.OrgId
		barcodes[i].MarketplaceID = item.MarketId
		barcodes[i].UpdatedAt = item.UpdateAt.ToTime()
	}
	return errors.Join(repository.SaveOrUpdateBarcodes(&barcodes))
}

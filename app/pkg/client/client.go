package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/yoda/app/pkg/mapper"
	"github.com/yoda/app/pkg/wbclient"
	"github.com/yoda/common/pkg/types"
	"log"
	"time"
)

type WBClient struct {
}

func NewWBClient() *WBClient {
	return &WBClient{}
}

func (c *WBClient) Parsing(listener *EventListener) error {
	apiKeyProvider, _ := securityprovider.NewSecurityProviderApiKey("header", "Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NJRCI6IjFiMzVmODljLTMyNGYtNGM3OS05NzhhLTkwMmYwODk3Mjc4YiJ9.WeYv1vqA46_9D5up2LRUeSBZCXxSBNcmH8lUhG9Jii0")
	clnt, err := wbclient.NewClientWithResponses("http://localhost:1080", wbclient.WithRequestEditorFn(apiKeyProvider.Intercept))
	if err != nil {
		return err
	}
	initTime, _ := time.Parse(time.DateOnly, "2020-01-01")
	dateFrom := wbclient.GetSupplierStocksParams{DateFrom: wbclient.DateFrom{
		Time: initTime,
	}}
	resp, err := clnt.GetSupplierStocksWithResponse(context.Background(), &dateFrom)
	if err != nil {
		return err
	}
	if resp.StatusCode()/100 != 2 {
		return errors.New(fmt.Sprintf("Http status: %s", resp.Status()))
	}
	items := *resp.JSON200
	if len(items) == 0 {
		log.Print("There aren't warehouses")
	}
	source := types.SourceTypeWB
	transactionId := (*listener).BeginOperation(&source)
	for index, item := range items {
		si, err := mapper.MapStockItem(&item)
		if err != nil {
			(*listener).EndOperation(transactionId, types.StatusTypeRejected)
			log.Printf("Couldn't map a value at row %d (%s)", index, err)
		}
		si.Source = &source
		si.Transaction = transactionId
		(*listener).WriteItem(si)
	}
	(*listener).EndOperation(transactionId, types.StatusTypeCompleted)
	return nil
}

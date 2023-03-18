package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"
)

const host = "http://fakehost"

type MockClient struct {
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	switch req.Method {
	case "GET":
		json := `[
        {
          "date": "2022-10-24T07:39:04",
          "lastChangeDate": "2022-11-01T04:40:19",
          "supplierArticle": "ИР039738",
          "techSize": "0",
          "barcode": "5411416057895",
          "totalPrice": 296,
          "discountPercent": 0,
          "isSupply": false,
          "isRealization": true,
          "promoCodeDiscount": 0,
          "warehouseName": "Электросталь",
          "countryName": "Россия",
          "oblastOkrugName": "",
          "regionName": "",
          "incomeID": 9012568,
          "saleID": "S3605371384",
          "odid": 600688604297,
          "spp": 0,
          "forPay": 245.68,
          "finishedPrice": 269,
          "priceWithDisc": 296,
          "nmId": 95755805,
          "subject": "Прокладки урологические",
          "category": "Здоровье",
          "brand": "Wellfix",
          "IsStorno": 0,
          "gNumber": "2121892670617299494",
          "sticker": "",
          "srid": "46761eefed344c6faf2e7e245f8a0ab5"
        }]`
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewReader([]byte(json))),
			Header:     http.Header{"Content-Type": []string{"application/json"}},
		}, nil
	case "POST":
		return nil, nil
	default:
		return nil, nil
	}
}

func TestApi(t *testing.T) {

	mockFunc := func(client *Client) error {
		client.Client = &MockClient{}
		return nil
	}

	client, err := NewClientWithResponses(host, mockFunc)
	if err != nil {
		t.Errorf("err: %v", err)
	}

	df := DateFrom{
		Time: time.Now(),
	}

	resp, err := client.GetWBSalesWithResponse(context.Background(), &GetWBSalesParams{
		DateFrom: df,
	})
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Errorf("status code: %v", resp.StatusCode())
	}
	items := *resp.JSON200
	if len(items) != 1 {
		t.Errorf("items len: %v", len(items))
	}

}

package api

import (
	"encoding/json"
	"testing"
)

func TestProductUnmarshal(t *testing.T) {
	jsonString := `{
                    "sku": 685064031,
                    "name": "Прокладки урологические Lady Ultra 28 штук Wellfix при недержании мочи",
                    "quantity": 1,
                    "offer_id": "ИР039738",
                    "price": "358.00",
                    "digital_codes": [],
                    "currency_code": "RUB"
                }`
	json.Unmarshal([]byte(jsonString), &PostingProduct{})
}

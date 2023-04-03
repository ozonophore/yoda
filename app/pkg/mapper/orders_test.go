package mapper

import (
	"github.com/stretchr/testify/assert"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/common/pkg/types"
	"testing"
	"time"
)

/**
* Function to test methdo MapOrder
 */
func TestMapOrder(t *testing.T) {
	transactionID := int64(3)
	source := generateOrder()
	order, err := MapOrder(&source, transactionID, "source", "ownerCode")
	assert.Nil(t, err)
	assert.Equal(t, transactionID, order.TransactionID)
	assert.Equal(t, "source", order.Source)
	assert.Equal(t, "2020-01-01", order.LastChangeDate.Format(time.DateOnly))
	assert.Equal(t, "10:30:00", order.LastChangeTime.Format(time.TimeOnly))
	assert.Equal(t, "2023-03-10", order.OrderDate.Format(time.DateOnly))
	assert.Equal(t, "11:20:00", order.OrderTime.Format(time.TimeOnly))
	assert.Equal(t, "supplierArticle", *order.SupplierArticle)
	assert.Equal(t, "techSize", *order.TechSize)
	assert.Equal(t, "barcode", *order.Barcode)
	assert.Equal(t, float64(12), order.DiscountPercent)
	assert.Equal(t, "warehouseName", *order.WarehouseName)
	assert.Equal(t, "oblast", *order.Oblast)
	assert.Equal(t, 43, *order.IncomeID)
	assert.Equal(t, int32(76), *order.Odid)
	assert.Equal(t, "gNumber", *order.GNumber)
	assert.Equal(t, "sticker", *order.Sticker)
	assert.Equal(t, "srid", *order.Srid)
	assert.Equal(t, "subject", *order.Subject)
	assert.Equal(t, "category", *order.Category)
	assert.Equal(t, "brand", *order.Brand)
	assert.Equal(t, "ownerCode", order.OwnerCode)
}

func generateOrder() api.OrdersItem {
	tm, _ := time.Parse(time.DateOnly+"T"+time.TimeOnly, "2020-01-01T10:30:00")
	lastChangeDate := types.CustomTime(tm)
	date := "2023-03-10T11:20:00"
	supplierArticle := "supplierArticle"
	techSize := "techSize"
	barcode := "barcode"
	discountPercent := 12
	warehouseName := "warehouseName"
	oblast := "oblast"
	IncomeID := 43
	odid := 76
	gNumber := "gNumber"
	sticker := "sticker"
	srid := "srid"
	subject := "subject"
	category := "category"
	brand := "brand"
	nmId := 12
	totalPrice := 12.50
	return api.OrdersItem{
		Date:            &date,
		LastChangeDate:  &lastChangeDate,
		SupplierArticle: &supplierArticle,
		TechSize:        &techSize,
		Barcode:         &barcode,
		DiscountPercent: &discountPercent,
		WarehouseName:   &warehouseName,
		Oblast:          &oblast,
		IncomeID:        &IncomeID,
		Odid:            &odid,
		GNumber:         &gNumber,
		Sticker:         &sticker,
		Srid:            &srid,
		Subject:         &subject,
		Category:        &category,
		Brand:           &brand,
		NmId:            &nmId,
		TotalPrice:      &totalPrice,
	}
}

package mapper

import (
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/utils"
	"strconv"
	"time"
)

// MapStockItem maps api.StocksItem to model.StockItem
func MapStockItem(s *api.StocksItem) (*model.StockItem, error) {
	externalCode := strconv.Itoa(*s.NmId)
	changeDate, err := time.Parse(time.DateOnly+"T"+time.TimeOnly, *s.LastChangeDate)
	if err != nil {
		return nil, err
	}
	qp := int32(0)
	return &model.StockItem{
		LastChangeDate:   changeDate,
		LastChangeTime:   changeDate,
		SupplierArticle:  s.SupplierArticle,
		Barcode:          s.Barcode,
		WarehouseName:    *s.WarehouseName,
		ExternalCode:     &externalCode,
		Subject:          s.Subject,
		Category:         s.Category,
		Brand:            s.Brand,
		Price:            utils.Float32ToFloat64(s.Price),
		Discount:         utils.Float32ToFloat64(s.Discount),
		Quantity:         *utils.IntToInt32(s.Quantity),
		QuantityFull:     *utils.IntToInt32(s.QuantityFull),
		QuantityPromised: &qp,
		SCCode:           s.SCCode,
		DaysOnSite:       utils.IntToInt32(s.DaysOnSite),
		IsRealization:    s.IsRealization,
		TechSize:         s.TechSize,
	}, nil
}

func MapRowItem(s *api.RowItem, d *time.Time) (*model.StockItem, error) {
	externalCode := strconv.Itoa(*s.Sku)
	q := int32(*s.FreeToSellAmount)
	qf := int32(*s.FreeToSellAmount) + int32(*s.ReservedAmount)
	qp := int32(*s.PromisedAmount)
	return &model.StockItem{
		LastChangeDate:   *d,
		LastChangeTime:   *d,
		SupplierArticle:  s.ItemCode,
		Barcode:          nil,
		WarehouseName:    *s.WarehouseName,
		ExternalCode:     &externalCode,
		Name:             s.ItemName,
		Category:         nil,
		Brand:            nil,
		Price:            nil,
		Discount:         nil,
		Quantity:         q,
		QuantityFull:     qf,
		QuantityPromised: &qp,
	}, nil
}

package mapper

import (
	"github.com/yoda/app/pkg/wbclient"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/utils"
	"strconv"
	"time"
)

func MapStockItem(s *wbclient.StocksItem) (*model.StockItem, error) {
	externalCode := strconv.Itoa(*s.NmId)
	changeDate, err := time.Parse(time.DateOnly+"T"+time.TimeOnly, *s.LastChangeDate)
	if err != nil {
		return nil, err
	}
	return &model.StockItem{
		LastChangeDate:  changeDate,
		LastChangeTime:  changeDate,
		SupplierArticle: s.SupplierArticle,
		Barcode:         s.Barcode,
		WarehouseName:   s.WarehouseName,
		ExternalCode:    &externalCode,
		Subject:         s.Subject,
		Category:        s.Category,
		Brand:           s.Brand,
		Price:           utils.Float32ToFloat64(s.Price),
		Discount:        utils.Float32ToFloat64(s.Discount),
		Quantity:        utils.IntToInt32(s.Quantity),
		QuantityFull:    utils.IntToInt32(s.QuantityFull),
	}, nil
}

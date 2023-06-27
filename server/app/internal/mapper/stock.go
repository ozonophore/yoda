package mapper

import (
	"github.com/yoda/app/internal/api"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
	"github.com/yoda/common/pkg/utils"
	"strconv"
	"strings"
	"time"
)

// MapStockItem maps api.StocksItem to model.StockItem
func MapStockItem(s *api.StocksItem) (*model.StockItem, error) {
	externalCode := strconv.Itoa(*s.NmId)
	changeDate := types.CustomTimeToTime(s.LastChangeDate)
	qp := int32(0)
	return &model.StockItem{
		TransactionDate:  time.Now(),
		LastChangeDate:   *changeDate,
		LastChangeTime:   *changeDate,
		SupplierArticle:  s.SupplierArticle,
		Barcode:          s.Barcode,
		WarehouseName:    strings.ToUpper(*s.WarehouseName),
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
		CardCreated:      changeDate.AddDate(0, 0, *s.DaysOnSite*-1),
	}, nil
}

func MapRowItem(s *api.RowItem, d *time.Time) (*model.StockItem, error) {
	externalCode := strconv.FormatInt(*s.Sku, 10)
	q := int32(*s.FreeToSellAmount)
	qf := int32(*s.FreeToSellAmount) + int32(*s.ReservedAmount)
	qp := int32(*s.PromisedAmount)
	return &model.StockItem{
		TransactionDate:  time.Now(),
		LastChangeDate:   *d,
		LastChangeTime:   *d,
		SupplierArticle:  s.ItemCode,
		Barcode:          nil,
		WarehouseName:    strings.ToUpper(*s.WarehouseName),
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

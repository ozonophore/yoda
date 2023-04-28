package mapper

import (
	"fmt"
	"github.com/yoda/app/pkg/api"
	"github.com/yoda/common/pkg/types"
	"github.com/yoda/common/pkg/utils"
	"testing"
	"time"
)

func TestMapStockItem(t *testing.T) {
	discount := float32(10)
	price := float32(100)
	scCode := "scCode"
	barcode := "barcode"
	category := "category"
	daysOnSite := 10
	isRealization := true
	isSupply := true
	lastChangeDate := "2020-11-01T12:30:02"
	nmId := 10
	quantity := 10
	quantityFull := 10
	subject := "subject"
	supplierArticle := "supplierArticle"
	techSize := "techSize"
	warehouseName := "warehouseName"

	s := api.StocksItem{
		Discount:        &discount,
		Price:           &price,
		SCCode:          &scCode,
		Barcode:         &barcode,
		Category:        &category,
		DaysOnSite:      &daysOnSite,
		IsRealization:   &isRealization,
		IsSupply:        &isSupply,
		LastChangeDate:  types.StringToCustomTime(lastChangeDate),
		NmId:            &nmId,
		Quantity:        &quantity,
		QuantityFull:    &quantityFull,
		Subject:         &subject,
		SupplierArticle: &supplierArticle,
		TechSize:        &techSize,
		WarehouseName:   &warehouseName,
	}
	r, _ := MapStockItem(&s)
	if *r.Discount != *utils.Float32ToFloat64(s.Discount) {
		t.Fatalf("Discount is not equal")
	}
	if *r.Price != *utils.Float32ToFloat64(s.Price) {
		t.Fatalf("Price is not equal")
	}
	if *r.SCCode != *s.SCCode {
		t.Fatalf("SCCode is not equal")
	}
	if *r.Barcode != *s.Barcode {
		t.Fatalf("Barcode is not equal")
	}
	if *r.Category != *s.Category {
		t.Fatalf("Category is not equal")
	}
	if *r.DaysOnSite != *utils.IntToInt32(s.DaysOnSite) {
		t.Fatalf("DaysOnSite is not equal")
	}
	if *r.IsRealization != *s.IsRealization {
		t.Fatalf("IsRealization is not equal")
	}
	lcd := r.LastChangeDate.Format(time.DateOnly + "T" + time.TimeOnly)
	if lcd != s.LastChangeDate.ToString() {
		t.Fatalf("LastChangeDate is not equal")
	}
	if *r.ExternalCode != fmt.Sprintf("%d", *s.NmId) {
		t.Fatalf("NmId is not equal")
	}
	if r.Quantity != *utils.IntToInt32(s.Quantity) {
		t.Fatalf("Quantity is not equal")
	}
	if r.QuantityFull != *utils.IntToInt32(s.QuantityFull) {
		t.Fatalf("QuantityFull is not equal")
	}
	if *r.Subject != *s.Subject {
		t.Fatalf("Subject is not equal")
	}
	if *r.SupplierArticle != *s.SupplierArticle {
		t.Fatalf("SupplierArticle is not equal")
	}
	if *r.TechSize != *s.TechSize {
		t.Fatalf("TechSize is not equal")
	}
	if r.WarehouseName != *s.WarehouseName {
		t.Fatalf("WarehouseName is not equal")
	}
}

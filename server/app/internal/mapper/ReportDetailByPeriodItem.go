package mapper

import (
	"github.com/yoda/app/internal/api"
	"github.com/yoda/common/pkg/model"
	"github.com/yoda/common/pkg/types"
)

func MapReportDetailByPeriodItem(data *api.ReportDetailByPeriodItem, transactionId int64, source, ownerCode string, barcodeId, itemId, message *string) *model.ReportDetailByPeriod {
	return &model.ReportDetailByPeriod{
		TransactionID:            transactionId,
		Source:                   source,
		OwnerCode:                ownerCode,
		AcquiringBank:            data.AcquiringBank,
		AcquiringFee:             data.AcquiringFee,
		AdditionalPayment:        data.AdditionalPayment,
		Barcode:                  data.Barcode,
		BonusTypeName:            data.BonusTypeName,
		BrandName:                data.BrandName,
		CommissionPercent:        data.CommissionPercent,
		CreateDt:                 types.CustomTimeToTime(data.CreateDt),
		DateFrom:                 types.CustomTimeToTime(data.DateFrom),
		DateTo:                   types.CustomTimeToTime(data.DateTo),
		DeclarationNumber:        data.DeclarationNumber,
		DeliveryAmount:           data.DeliveryAmount,
		DeliveryRub:              data.DeliveryRub,
		DocNumber:                data.DocNumber,
		DocTypeName:              data.DocTypeName,
		GiBoxTypeName:            data.GiBoxTypeName,
		GiID:                     data.GiId,
		Kiz:                      data.Kiz,
		NmID:                     data.NmId,
		OfficeName:               data.OfficeName,
		OrderDt:                  types.CustomTimeToTime(data.OrderDt),
		Penalty:                  data.Penalty,
		PpvzForPay:               data.PpvzForPay,
		PpvzInn:                  data.PpvzInn,
		PpvzKvwPrc:               data.PpvzKvwPrc,
		PpvzKvwPrcBase:           data.PpvzKvwPrcBase,
		PpvzOfficeID:             data.PpvzOfficeId,
		PpvzOfficeName:           data.PpvzOfficeName,
		PpvzReward:               data.PpvzReward,
		PpvzSalesCommission:      data.PpvzSalesCommission,
		PpvzSppPrc:               data.PpvzSppPrc,
		PpvzSupplierID:           data.PpvzSupplierId,
		PpvzSupplierName:         data.PpvzSupplierName,
		PpvzVw:                   data.PpvzVw,
		PpvzVwNds:                data.PpvzVwNds,
		ProductDiscountForReport: data.ProductDiscountForReport,
		Quantity:                 data.Quantity,
		RealizationreportID:      data.RealizationreportId,
		RetailAmount:             data.RetailAmount,
		RetailPriceWithdiscRub:   data.RetailPriceWithdiscRub,
		ReturnAmount:             data.ReturnAmount,
		RetailPrice:              data.RetailPrice,
		Rid:                      data.Rid,
		RrDt:                     types.CustomTimeToTime(data.RrDt),
		RrdID:                    data.RrdId,
		SaName:                   data.SaName,
		SaleDt:                   types.CustomTimeToTime(data.SaleDt),
		SalePercent:              data.SalePercent,
		ShkID:                    data.ShkId,
		SiteCountry:              data.SiteCountry,
		Srid:                     data.Srid,
		StickerID:                data.StickerId,
		SubjectName:              data.SubjectName,
		SupplierOperName:         data.SupplierOperName,
		SupplierPromo:            data.SupplierPromo,
		SuppliercontractCode:     data.SuppliercontractCode,
		TsName:                   data.TsName,
		BarcodeID:                barcodeId,
		ItemID:                   itemId,
		Message:                  message,
	}
}
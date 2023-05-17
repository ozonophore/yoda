package model

type OrgSrcBarcode struct {
	OrgCode   string `gorm:"column:org_code" json:"org_code"`
	Source    string `gorm:"column:source" json:"source"`
	Barcode   string `gorm:"column:barcode" json:"barcode"`
	BarcodeId string `gorm:"column:barcode_id" json:"barcode_id"`
	ItemId    string `gorm:"column:item_id" json:"item_id"`
	Name      string `gorm:"column:name" json:"name"`
}

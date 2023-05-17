package dictionary

import "sync"

var instance *ItemDecoder = newItemDecoder()

type ItemNotFoundError struct {
	org     string
	src     string
	barcode string
}

func NewNotFoundError(org, src, barcode string) *ItemNotFoundError {
	return &ItemNotFoundError{
		org:     org,
		src:     src,
		barcode: barcode,
	}
}

func (i *ItemNotFoundError) Error() string {
	return `NotFoundException: Item for org ` + i.org + ` source ` + i.src + ` barcode ` + i.barcode + ` not found`
}

type keyItem struct {
	OrgCode string
	Source  string
	Barcode string
}

type ItemBarcode struct {
	BarcodeId string
	ItemId    string
}

type ItemDecoder struct {
	storage map[keyItem]*ItemBarcode
	mutex   sync.Mutex
}

func newItemDecoder() *ItemDecoder {
	return &ItemDecoder{
		storage: make(map[keyItem]*ItemBarcode, 1024),
		mutex:   sync.Mutex{},
	}
}

func GetItemDecoder() *ItemDecoder {
	return instance
}

func (d *ItemDecoder) Add(orgCode, source, barcode, barcodeId, itemId string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	key := keyItem{
		OrgCode: orgCode,
		Source:  source,
		Barcode: barcode,
	}
	_, ok := d.storage[key]
	if ok {
		return
	}
	d.storage[key] = &ItemBarcode{
		BarcodeId: barcodeId,
		ItemId:    itemId,
	}
}

func (d *ItemDecoder) Decode(orgCode, source, barcode string) (*ItemBarcode, error) {
	key := keyItem{
		OrgCode: orgCode,
		Source:  source,
		Barcode: barcode,
	}
	value := d.storage[key]
	if value != nil {
		return value, nil
	}
	return nil, NewNotFoundError(orgCode, source, barcode)
}

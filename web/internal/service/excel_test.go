package service

import (
	"github.com/yoda/web/internal/storage"
	"testing"
)

func TestGetExcelHeaders(t *testing.T) {
	stock := storage.Stock{}
	GetExcelHeaders(stock)
}

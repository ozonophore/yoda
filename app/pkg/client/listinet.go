package client

import "github.com/yoda/common/pkg/model"

type EventListener interface {
	BeginOperation(t *string) *int64                // performs operation starting
	EndOperation(transaction *int64, status string) // performs write item
	WriteItem(model *model.StockItem)               // performs operation ending
}

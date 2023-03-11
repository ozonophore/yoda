package service

import (
	"testing"
)

func TestCallbackBatch(t *testing.T) {
	type customType struct {
		Id int
	}
	items := []customType{{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}, {11}, {12}, {13}}

	batchSize := 3

	var newItems []customType
	callbackFunc := func(items *[]customType) error {
		for _, item := range *items {
			newItems = append(newItems, item)
		}
		return nil
	}
	CallbackBatch[customType](&items, batchSize, callbackFunc)
	if len(newItems) != len(items) {
		t.Errorf("len(newItems): %d, len(items): %d", len(newItems), len(items))
	}
	if items[0] != newItems[0] {
		t.Errorf("items[0]: %d, newItems[0]: %d", items[0], newItems[0])
	}
	if items[1] != newItems[1] {
		t.Errorf("items[1]: %d, newItems[1]: %d", items[1], newItems[1])
	}
	if items[2] != newItems[2] {
		t.Errorf("items[2]: %d, newItems[2]: %d", items[2], newItems[2])
	}
	if items[3] != newItems[3] {
		t.Errorf("items[3]: %d, newItems[3]: %d", items[3], newItems[3])
	}
}

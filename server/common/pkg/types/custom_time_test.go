package types

import (
	"encoding/json"
	"testing"
)

type testStructure struct {
	Date CustomTime `json:"date"`
}

func TestCustomTime_UnmarshalJSON(t *testing.T) {
	source := `{"Date":"2023-05-08T00:00:00Z"}`
	data := []byte(source)
	var ct testStructure
	err := json.Unmarshal(data, &ct)
	if err != nil {
		t.Error(err)
	}
}

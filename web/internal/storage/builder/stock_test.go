package builder

import (
	"testing"
	"time"
)

func TestStockBuilder(t *testing.T) {
	builder := NewStockSQLBuilder(time.Now())

	builder.Sources(&[]string{"1", "2"})
	limit := 10
	offset := 20
	builder.Limit(&limit)
	builder.Offset(&offset)
	filter := "filter"
	builder.Filter(&filter)
	sql, _ := builder.Build()
	print(sql)
}

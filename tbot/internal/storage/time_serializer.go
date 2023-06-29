package storage

import (
	"context"
	"gorm.io/gorm/schema"
	"reflect"
	"time"
)

type TimeSerializer struct {
}

func (TimeSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	fieldValue := reflect.New(field.FieldType)

	if dbValue != nil {
		time, err := time.Parse(time.TimeOnly, dbValue.(string))
		if err != nil {
			return err
		}
		fieldValue.Elem().Set(reflect.ValueOf(time))

		err = nil
	}

	field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
	return
}

// Value implements serializer interface
func (TimeSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return fieldValue, nil
}

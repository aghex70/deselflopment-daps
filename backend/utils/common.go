package common

import (
	"github.com/satori/go.uuid"
	"reflect"
)

func GenerateUUID() string {
	u := uuid.NewV4()
	return u.String()
}

func StructToMap(data interface{}, ignoredField string) map[string]interface{} {
	result := make(map[string]interface{})

	v := reflect.ValueOf(data)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")

		if tag == "" || tag == ignoredField {
			continue
		}
		fieldValue := v.Field(i)

		// Handle different kinds of field values
		switch fieldValue.Kind() {
		case reflect.Ptr:
			if fieldValue.IsNil() {
				continue
			}
			fieldValue = fieldValue.Elem()
			result[tag] = fieldValue.Interface()
		default:
			result[tag] = fieldValue.Interface()
		}
	}
	return result
}

package validations

import "reflect"

func setErrorMessage(v any, fieldName, msg string) {
	if v == nil {
		return
	}
	switch t := v.(type) {
	case map[string]string:
		t[fieldName] = msg
	default:
		structVal := reflect.ValueOf(v)
		if structVal.Kind() != reflect.Ptr || structVal.IsNil() {
			return
		}
		structVal = structVal.Elem()
		field := structVal.FieldByName(fieldName)
		field.Set(reflect.ValueOf(msg))
	}
}

func getFieldValueByName(v any, name string) any {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return nil
	}
	fieldVal := val.FieldByName(name)
	if !fieldVal.IsValid() {
		return nil
	}
	return fieldVal.Interface()
}

package toolz

import (
	"fmt"
	"log/slog"
	"reflect"
	"time"
)

// Slogify uses reflection to recursively traverse a value and build a slog.Attr
func Slogify(name string, obj any) slog.Attr {
	val := reflect.ValueOf(obj)

	// If it's a pointer, get the underlying value
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Bool:
		return slog.Bool(name, val.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return slog.Int64(name, val.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return slog.Uint64(name, val.Uint())
	case reflect.Float32, reflect.Float64:
		return slog.Float64(name, val.Float())
	case reflect.String:
		return slog.String(name, val.String())
	case reflect.Struct:
		if val.Type() == reflect.TypeOf(time.Time{}) {
			return slog.Time(name, val.Interface().(time.Time))
		}
		if val.Type() == reflect.TypeOf(time.Duration(0)) {
			return slog.Duration(name, val.Interface().(time.Duration))
		}

		// For other structs, create a nested group
		attrs := make([]any, 0, val.NumField())
		t := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := t.Field(i)
			if !fieldType.IsExported() {
				continue
			}
			attrs = append(attrs, Slogify(fieldType.Name, field.Interface()))
		}
		return slog.Group(name, attrs...)

	case reflect.Slice, reflect.Array:
		attrs := make([]any, 0, val.Len())
		for i := 0; i < val.Len(); i++ {
			attrs = append(attrs, Slogify(fmt.Sprintf("%d", i), val.Index(i).Interface()))
		}
		return slog.Group(name, attrs...)

	case reflect.Map:
		attrs := make([]any, 0, val.Len())
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			attrs = append(attrs, Slogify(fmt.Sprint(k.Interface()), iter.Value().Interface()))
		}
		return slog.Group(name, attrs...)

	default:
		return slog.String(name, fmt.Sprintf("%v", val.Interface()))
	}
}

package toolz

import (
	"log/slog"
	"reflect"
)

// The AllEmpty function checks if all pointers are nil and all
// slices, maps and arrays are empty.
func AllEmpty(pointers ...any) bool {
	for _, p := range pointers {
		// untyped
		if p == nil {
			continue
		}

		// typed
		k := reflect.TypeOf(p).Kind()
		switch k {
		case reflect.Ptr:
			s := reflect.ValueOf(p)
			if !s.IsNil() {
				return false
			}
		case reflect.Func:
			f := reflect.ValueOf(p)
			if !f.IsNil() {
				return false
			}
		case reflect.Chan:
			c := reflect.ValueOf(p)
			if !c.IsNil() {
				return false
			}
		case reflect.Interface:
			i := reflect.ValueOf(p)
			if !i.IsNil() {
				return false
			}
		case reflect.Slice:
			s := reflect.ValueOf(p)
			if s.Len() > 0 {
				return false
			}
		case reflect.Map:
			m := reflect.ValueOf(p)
			if m.Len() > 0 {
				return false
			}
		case reflect.Array:
			a := reflect.ValueOf(p)
			if a.Len() > 0 {
				return false
			}
		default:
			slog.Warn("AllEmpty: unknown type", "type", k)
			return false
		}
	}
	return true
}

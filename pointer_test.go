package toolz

import (
	"testing"
)

func TestP(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		val := 42
		ptr := P(val)
		if *ptr != val {
			t.Errorf("expected %d, got %d", val, *ptr)
		}
	})

	t.Run("string", func(t *testing.T) {
		val := "hello"
		ptr := P(val)
		if *ptr != val {
			t.Errorf("expected %s, got %s", val, *ptr)
		}
	})

	t.Run("float64", func(t *testing.T) {
		val := 3.14
		ptr := P(val)
		if *ptr != val {
			t.Errorf("expected %f, got %f", val, *ptr)
		}
	})

	t.Run("struct", func(t *testing.T) {
		type example struct {
			Field1 int
			Field2 string
		}
		val := example{Field1: 1, Field2: "test"}
		ptr := P(val)
		if *ptr != val {
			t.Errorf("expected %+v, got %+v", val, *ptr)
		}
	})

	t.Run("nil pointer", func(t *testing.T) {
		var val *int
		ptr := P(val)
		if *ptr != val {
			t.Errorf("expected %v, got %v", val, *ptr)
		}
	})
}

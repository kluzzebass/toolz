package toolz

import (
	"testing"
)

// test the AllNil function
func TestAllEmpty(t *testing.T) {
	// test cases

	// nil channel
	var c1 chan int

	// non-nil channel
	c2 := make(chan int)

	// nil function
	var f1 func()

	// non-nil function
	f2 := func() {}

	// empty array
	var a1 []int

	// non-empty array
	a2 := []int{1}

	// empty interface
	var i1 interface{}

	// non-empty interface
	i2 := 1

	cases := []struct {
		name     string
		pointers any
		expected bool
	}{
		{"untyped", nil, true}, // untyped nil
		{"interface slice", []interface{}{nil, nil, nil}, false},
		{"interface slice", []interface{}{}, true},
		{"int slice", []int{}, true},
		{"string slice", []string{}, true},
		{"slice", []interface{}{nil, []int{}, nil}, false},
		{"slice", []interface{}{nil, []int{1}, nil}, false},
		{"channel", c1, true},
		{"channel", c2, false},
		{"function", f1, true},
		{"function", f2, false},
		{"map", map[string]int{}, true},
		{"map", map[string]int{"blah": 2}, false},
		{"array", a1, true},
		{"array", a2, false},
		{"interface", i1, true},
		{"interface", i2, false},
	}

	// run test cases
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := AllEmpty(c.pointers)
			if actual != c.expected {
				t.Errorf("expected %t, got %t", c.expected, actual)
			}
		})
	}
}

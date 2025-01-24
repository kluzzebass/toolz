package toolz

// Coalesce returns the first non-nil value in the list as a value of type T,
// and nil if all values are nil.
func Coalesce[T any](values ...*T) *T {
	for _, v := range values {
		if v != nil {
			return v
		}
	}
	return nil
}

// Coalesce returns the first non-nil value in the list as a value of type T,
// and a zero value if all values are nil.
func CoalesceOrZero[T any](values ...*T) T {
	if v := Coalesce(values...); v != nil {
		return *v
	}
	return *new(T)
}

package toolz

// P is a helper function to return a pointer to a value.
func P[T any](val T) *T {
	return &val
}

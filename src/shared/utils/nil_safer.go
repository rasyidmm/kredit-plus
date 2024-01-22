package utils

// ValueOrZero check if pointer of value is not nil and return its value if exists, otherwise return zero value of its type
func ValueOrZero[T any](v *T) T {
	if v == nil {
		var zeroVal T
		return zeroVal
	}

	return *v
}

// ValueOrDefault check if pointer of value is not nil and return its value if exists, otherwise return default value
func ValueOrDefault[T any](v *T, d T) T {
	if v == nil {
		return d
	}

	return *v
}

// ValueToPtr convert value to pointer address of value
func ValueToPtr[T any](v T) *T {
	return &v
}

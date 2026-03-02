package converter

// ToPointer нь утгын pointer-ийг буцаана.
func ToPointer[T any](v T) *T {
	return &v
}

// FromPointer нь pointer-ийн утгыг буцаана. nil бол zero value буцаана.
func FromPointer[T any](p *T) T {
	if p == nil {
		var zero T
		return zero
	}
	return *p
}

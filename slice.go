package slice

type Slice[T comparable] struct {
	inner []T
}

func From[T comparable](from []T) *Slice[T] {
	return &Slice[T]{inner: from}
}

func (slc Slice[T]) Into() []T {
	return slc.inner
}

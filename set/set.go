package set

type Set[T any] struct {
	items    []T
	iterator int
}

// New creates new empty collection
func New[T any]() *Set[T] {
	return &Set[T]{
		items: make([]T, 0),
	}
}

// NewWithSize creates new empty collection with given capacity
func NewWithSize[T any](size int) *Set[T] {
	return &Set[T]{
		items: make([]T, 0, size),
	}
}

// NewWithList creates new collection with given slice with values
func NewWithList[T any](items ...T) *Set[T] {
	return &Set[T]{
		items: items,
	}
}

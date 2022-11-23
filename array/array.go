package array

type Array[T any] struct {
	items    []T
	iterator int
}

// New creates new empty collection
func New[T any]() *Array[T] {
	return &Array[T]{
		items: make([]T, 0),
	}
}

// NewWithSize creates new empty collection with given capacity
func NewWithSize[T any](size int) *Array[T] {
	return &Array[T]{
		items: make([]T, 0, size),
	}
}

// NewWithList creates new collection with given slice with values
func NewWithList[T any](items ...T) *Array[T] {
	return &Array[T]{
		items: items,
	}
}

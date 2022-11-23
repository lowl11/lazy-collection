package type_list

type TypeList[T, X any] struct {
	items    []T
	iterator int
}

// New creates new empty collection
func New[T, X any]() *TypeList[T, X] {
	return &TypeList[T, X]{
		items: make([]T, 0),
	}
}

// NewWithSize creates new empty collection with given capacity
func NewWithSize[T, X any](size int) *TypeList[T, X] {
	return &TypeList[T, X]{
		items: make([]T, 0, size),
	}
}

// NewWithList creates new collection with given slice with values
func NewWithList[T, X any](items ...T) *TypeList[T, X] {
	return &TypeList[T, X]{
		items: items,
	}
}

package array

type Array[T any] struct {
	items    []T
	iterator int
}

func New[T any]() *Array[T] {
	return &Array[T]{
		items: make([]T, 0),
	}
}

func NewWithSize[T any](size int) *Array[T] {
	return &Array[T]{
		items: make([]T, 0, size),
	}
}

func NewWithList[T any](items ...T) *Array[T] {
	return &Array[T]{
		items: items,
	}
}

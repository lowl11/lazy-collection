package set

type Set[T any] struct {
	items    []T
	iterator int
}

func New[T any]() *Set[T] {
	return &Set[T]{
		items: make([]T, 0),
	}
}

func NewWithSize[T any](size int) *Set[T] {
	return &Set[T]{
		items: make([]T, 0, size),
	}
}

func NewWithList[T any](items ...T) *Set[T] {
	return &Set[T]{
		items: items,
	}
}

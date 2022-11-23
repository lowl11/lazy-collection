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

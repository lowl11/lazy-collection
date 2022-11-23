package array

type Array[T any] struct {
	items []T
}

func New[T any]() *Array[T] {
	return &Array[T]{
		items: make([]T, 0),
	}
}

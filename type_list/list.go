package type_list

import "github.com/lowl11/lazy-collection/helper"

const (
	defaultCapacity  = 100
	increaseCapacity = 2
)

type TypeList[T, X any] struct {
	items    []T
	iterator int
}

// New creates new empty collection
func New[T, X any]() *TypeList[T, X] {
	return &TypeList[T, X]{
		items: make([]T, 0, defaultCapacity),
	}
}

// NewWithSize creates new empty collection with given capacity
func NewWithSize[T, X any](size int) *TypeList[T, X] {
	return &TypeList[T, X]{
		items: make([]T, 0, size*increaseCapacity),
	}
}

// NewWithList creates new collection with given slice with values
func NewWithList[T, X any](items ...T) *TypeList[T, X] {
	newList := make([]T, len(items), (defaultCapacity+len(items))*increaseCapacity)
	helper.Fill[T](newList, items)

	return &TypeList[T, X]{
		items: newList,
	}
}

package set

import "github.com/lowl11/lazy-collection/helper"

const (
	defaultCapacity  = 100
	increaseCapacity = 2
)

type Set[T any] struct {
	items    []T
	iterator int
}

// New creates new empty collection
func New[T any]() *Set[T] {
	return &Set[T]{
		items: make([]T, 0, defaultCapacity),
	}
}

// NewWithSize creates new empty collection with given capacity
func NewWithSize[T any](size int) *Set[T] {
	return &Set[T]{
		items: make([]T, 0, size*increaseCapacity),
	}
}

// NewWithList creates new collection with given slice with values
func NewWithList[T any](items ...T) *Set[T] {
	newList := make([]T, len(items), (defaultCapacity+len(items))*increaseCapacity)
	helper.Fill[T](newList, items)

	return &Set[T]{
		items: newList,
	}
}

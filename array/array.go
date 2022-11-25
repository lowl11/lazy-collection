package array

import (
	"github.com/lowl11/lazy-collection/helper"
)

const (
	defaultCapacity  = 100
	increaseCapacity = 2
)

type Array[T any] struct {
	items    []T
	iterator int
}

// New creates new empty collection
func New[T any]() *Array[T] {
	return &Array[T]{
		items: make([]T, 0, defaultCapacity),
	}
}

// NewWithSize creates new empty collection with given capacity
func NewWithSize[T any](size int) *Array[T] {
	return &Array[T]{
		items: make([]T, 0, size*increaseCapacity),
	}
}

// NewWithList creates new collection with given slice with values
func NewWithList[T any](items ...T) *Array[T] {
	newList := make([]T, len(items), (defaultCapacity+len(items))*increaseCapacity)
	helper.Fill[T](newList, items)

	return &Array[T]{
		items: newList,
	}
}

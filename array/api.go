package array

import (
	"sort"
)

// Push add element to the end. Add only if value is unique
func (array *Array[T]) Push(value T) *Array[T] {
	array.items = append(array.items, value)
	return array
}

// PushForward add element to the start
func (array *Array[T]) PushForward(value T) *Array[T] {
	array.items = append([]T{value}, array.items...)
	return array
}

// PushList add slice/array of elements to the end
func (array *Array[T]) PushList(items ...T) *Array[T] {
	array.items = append(array.items, items...)
	return array
}

// Pop remove last element
func (array *Array[T]) Pop() *Array[T] {
	size := array.Size()
	if size == 0 {
		return array
	}

	array.items = array.items[:size-1]
	return array
}

// PopForward remove first element
func (array *Array[T]) PopForward() *Array[T] {
	size := array.Size()
	if size == 0 {
		return array
	}

	array.items = array.items[1:]
	return array
}

// Remove remove element on index
func (array *Array[T]) Remove(index int) *Array[T] {
	if array.Empty() || array.indexOut(index) {
		return array
	}

	array.items = append(array.items[:index], array.items[index+1:]...)
	return array
}

// RemoveAll remove all elements
func (array *Array[T]) RemoveAll() *Array[T] {
	array.items = make([]T, 0)
	return array
}

// Clear remove all elements with saving capacity (previous size)
func (array *Array[T]) Clear() *Array[T] {
	array.items = make([]T, 0, array.Size())
	return array
}

// Get returns element on index
func (array *Array[T]) Get(index int) T {
	size := array.Size()

	if index > size || size == 0 {
		emptyItem := new(T)
		return *emptyItem
	}

	return array.items[index]
}

// Empty check is empty or not
func (array *Array[T]) Empty() bool {
	return array.Size() == 0
}

// Size returns size of collection
func (array *Array[T]) Size() int {
	return len(array.items)
}

// Sort implementation of sort package
func (array *Array[T]) Sort(compareFunc func(i, j int) bool) *Array[T] {
	sort.Slice(array.items, compareFunc)
	return array
}

// Next moves iterator
func (array *Array[T]) Next() bool {
	iterator := array.iterator
	size := array.Size()

	if size == 0 || iterator == size {
		return false
	}

	array.iterator++
	return true
}

// Value returns value on iterator
func (array *Array[T]) Value() T {
	size := array.Size()
	if size == 0 {
		return *new(T)
	}

	return array.items[array.iterator-1]
}

// Reset set iterator to the start
func (array *Array[T]) Reset() *Array[T] {
	array.iterator = 0
	return array
}

// Clone create new collection with the same data
func (array *Array[T]) Clone() *Array[T] {
	return NewWithList[T](array.items...)
}

// CloneRange clone collection but with range
func (array *Array[T]) CloneRange(from, to int) *Array[T] {
	if from < 0 {
		from = 0
	}

	if array.indexOut(to) {
		to = array.Size()
	}

	return NewWithList[T](array.SliceRange(from, to)...)
}

// IndexOf search element and returns index
func (array *Array[T]) IndexOf(item T) int {
	return array.find(item)
}

// Slice returns slice of collection
func (array *Array[T]) Slice() []T {
	return array.items
}

// SliceRange returns slice of collection with range
func (array *Array[T]) SliceRange(from, to int) []T {
	if from < 0 {
		from = 0
	}

	if array.indexOut(to) {
		to = array.Size()
	}

	return array.items[from:to]
}

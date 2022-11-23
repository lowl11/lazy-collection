package array

import (
	"sort"
)

func (array *Array[T]) Push(value T) *Array[T] {
	array.items = append(array.items, value)
	return array
}

func (array *Array[T]) PushForward(value T) *Array[T] {
	array.items = append([]T{value}, array.items...)
	return array
}

func (array *Array[T]) PushList(items ...T) *Array[T] {
	array.items = append(array.items, items...)
	return array
}

func (array *Array[T]) Pop() *Array[T] {
	size := array.Size()
	if size == 0 {
		return array
	}

	array.items = array.items[:size-1]
	return array
}

func (array *Array[T]) PopForward() *Array[T] {
	size := array.Size()
	if size == 0 {
		return array
	}

	array.items = array.items[1:]
	return array
}

func (array *Array[T]) Remove(index int) *Array[T] {
	if array.Empty() || array.indexOut(index) {
		return array
	}

	array.items = append(array.items[:index], array.items[index+1:]...)
	return array
}

func (array *Array[T]) RemoveAll() *Array[T] {
	array.items = make([]T, 0)
	return array
}

func (array *Array[T]) Clear() *Array[T] {
	array.items = make([]T, 0, array.Size())
	return array
}

func (array *Array[T]) Get(index int) T {
	size := array.Size()

	if index > size || size == 0 {
		emptyItem := new(T)
		return *emptyItem
	}

	return array.items[index]
}

func (array *Array[T]) Empty() bool {
	return array.Size() == 0
}

func (array *Array[T]) Size() int {
	return len(array.items)
}

func (array *Array[T]) Sort(compareFunc func(i, j int) bool) *Array[T] {
	sort.Slice(array.items, compareFunc)
	return array
}

func (array *Array[T]) Next() bool {
	iterator := array.iterator
	size := array.Size()

	if size == 0 || iterator == size {
		return false
	}

	array.iterator++
	return true
}

func (array *Array[T]) Value() T {
	size := array.Size()
	if size == 0 {
		return *new(T)
	}

	return array.items[array.iterator-1]
}

func (array *Array[T]) Reset() *Array[T] {
	array.iterator = 0
	return array
}

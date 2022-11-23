package array

import "sort"

func (array *Array[T]) Push(value T) *Array[T] {
	array.items = append(array.items, value)
	return array
}

func (array *Array[T]) PushForward(value T) *Array[T] {
	array.items = append([]T{value}, array.items...)
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

func (array *Array[T]) Get(index int) T {
	size := array.Size()

	if index > size || size == 0 {
		emptyItem := new(T)
		return *emptyItem
	}

	return array.items[index]
}

func (array *Array[T]) Size() int {
	return len(array.items)
}

func (array *Array[T]) Sort(compareFunc func(i, j int) bool) *Array[T] {
	sort.Slice(array.items, compareFunc)
	return array
}

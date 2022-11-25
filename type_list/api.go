package type_list

import (
	"github.com/lowl11/lazy-collection/array"
	"sort"
)

// Push add element to the end. Add only if value is unique
func (typeList *TypeList[T, X]) Push(value T) *TypeList[T, X] {
	typeList.items = append(typeList.items, value)
	return typeList
}

// PushForward add element to the start
func (typeList *TypeList[T, X]) PushForward(value T) *TypeList[T, X] {
	typeList.items = append([]T{value}, typeList.items...)
	return typeList
}

// PushList add slice/array of elements to the end
func (typeList *TypeList[T, X]) PushList(items ...T) *TypeList[T, X] {
	typeList.items = append(typeList.items, items...)
	return typeList
}

// Pop remove last element
func (typeList *TypeList[T, X]) Pop() *TypeList[T, X] {
	size := typeList.Size()
	if size == 0 {
		return typeList
	}

	typeList.items = typeList.items[:size-1]
	return typeList
}

// PopForward remove first element
func (typeList *TypeList[T, X]) PopForward() *TypeList[T, X] {
	size := typeList.Size()
	if size == 0 {
		return typeList
	}

	typeList.items = typeList.items[1:]
	return typeList
}

// Set add element on index on the list
func (typeList *TypeList[T, X]) Set(index int, value T) *TypeList[T, X] {
	if typeList.indexOut(index) {
		return typeList
	}

	typeList.items[index] = value
	return typeList
}

// Remove remove element on index
func (typeList *TypeList[T, X]) Remove(index int) *TypeList[T, X] {
	if typeList.Empty() || typeList.indexOut(index) {
		return typeList
	}

	typeList.items = append(typeList.items[:index], typeList.items[index+1:]...)
	return typeList
}

// RemoveAll remove all elements
func (typeList *TypeList[T, X]) RemoveAll() *TypeList[T, X] {
	typeList.items = make([]T, 0)
	return typeList
}

// Clear remove all elements with saving capacity (previous size)
func (typeList *TypeList[T, X]) Clear() *TypeList[T, X] {
	typeList.items = make([]T, 0, typeList.Size())
	return typeList
}

// Get returns element on index
func (typeList *TypeList[T, X]) Get(index int) T {
	size := typeList.Size()

	if index > size || size == 0 {
		emptyItem := new(T)
		return *emptyItem
	}

	return typeList.items[index]
}

// Empty check is empty or not
func (typeList *TypeList[T, X]) Empty() bool {
	return typeList.Size() == 0
}

// Size returns size of collection
func (typeList *TypeList[T, X]) Size() int {
	return len(typeList.items)
}

// Sort implementation of sort package
func (typeList *TypeList[T, X]) Sort(compareFunc func(i, j int) bool) *TypeList[T, X] {
	sort.Slice(typeList.items, compareFunc)
	return typeList
}

// CopyTo copy collection to given
func (typeList *TypeList[T, X]) CopyTo(copyArray *TypeList[T, X]) *TypeList[T, X] {
	copyArray = NewWithList[T, X](typeList.items...)
	return typeList
}

// Reverse make reverse collection
func (typeList *TypeList[T, X]) Reverse() *TypeList[T, X] {
	for i, j := 0, len(typeList.items)-1; i < j; i, j = i+1, j-1 {
		typeList.items[i], typeList.items[j] = typeList.items[j], typeList.items[i]
	}

	return typeList
}

// Next moves iterator
func (typeList *TypeList[T, X]) Next() bool {
	iterator := typeList.iterator
	size := typeList.Size()

	if size == 0 || iterator == size {
		return false
	}

	typeList.iterator++
	return true
}

// Value returns value on iterator
func (typeList *TypeList[T, X]) Value() T {
	size := typeList.Size()
	if size == 0 {
		return *new(T)
	}

	return typeList.items[typeList.iterator-1]
}

// Reset set iterator to the start
func (typeList *TypeList[T, X]) Reset() *TypeList[T, X] {
	typeList.iterator = 0
	return typeList
}

// Clone create new collection with the same data
func (typeList *TypeList[T, X]) Clone() *TypeList[T, X] {
	return NewWithList[T, X](typeList.items...)
}

// CloneRange clone collection but with range
func (typeList *TypeList[T, X]) CloneRange(from, to int) *TypeList[T, X] {
	if from < 0 {
		from = 0
	}

	if typeList.indexOut(to) {
		to = typeList.Size()
	}

	return NewWithList[T, X](typeList.SliceRange(from, to)...)
}

// IndexOf search element and returns index
func (typeList *TypeList[T, X]) IndexOf(item T) int {
	return typeList.find(item)
}

// Contains true if collection owns element
func (typeList *TypeList[T, X]) Contains(item T) bool {
	return typeList.find(item) != -1
}

// Slice returns slice of collection
func (typeList *TypeList[T, X]) Slice() []T {
	return typeList.items
}

// SliceRange returns slice of collection with range
func (typeList *TypeList[T, X]) SliceRange(from, to int) []T {
	if from < 0 {
		from = 0
	}

	if typeList.indexOut(to) {
		to = typeList.Size()
	}

	return typeList.items[from:to]
}

// Select parse collection to another one with similar data type
func (typeList *TypeList[T, X]) Select(selectFunc func(sourceItem T) X) *array.Array[X] {
	newCollection := array.NewWithSize[X](typeList.Size())
	for _, item := range typeList.items {
		newCollection.Push(selectFunc(item))
	}
	return newCollection
}

// Sub returns new collection with range from this one
func (typeList *TypeList[T, X]) Sub(from, to int) *TypeList[T, X] {
	if from < 0 {
		from = 0
	}

	if typeList.indexOut(to) {
		to = typeList.Size()
	}

	return NewWithList[T, X](typeList.SliceRange(from, to)...)
}

package set

import (
	"sort"
)

// Push add element to the end
func (set *Set[T]) Push(value T) *Set[T] {
	if set.Contains(value) {
		return set
	}

	set.items = append(set.items, value)
	return set
}

// PushForward add element to the start
func (set *Set[T]) PushForward(value T) *Set[T] {
	if set.Contains(value) {
		return set
	}

	set.items = append([]T{value}, set.items...)
	return set
}

// PushList add slice/array of elements to the end
func (set *Set[T]) PushList(items ...T) *Set[T] {
	for _, item := range items {
		if !set.Contains(item) {
			set.items = append(set.items, item)
		}
	}

	return set
}

// Pop remove last element
func (set *Set[T]) Pop() *Set[T] {
	if set.Empty() {
		return set
	}

	set.items = set.items[:set.Size()-1]
	return set
}

// PopForward remove first element
func (set *Set[T]) PopForward() *Set[T] {
	if set.Empty() {
		return set
	}

	set.items = set.items[1:]
	return set
}

// Set add element on index on the list
func (set *Set[T]) Set(index int, value T) *Set[T] {
	if set.indexOut(index) {
		return set
	}

	set.items[index] = value
	return set
}

// Remove remove element on index
func (set *Set[T]) Remove(index int) *Set[T] {
	if set.Empty() || set.indexOut(index) {
		return set
	}

	set.items = append(set.items[:index], set.items[index+1:]...)
	return set
}

// RemoveAll remove all elements
func (set *Set[T]) RemoveAll() *Set[T] {
	set.items = make([]T, 0)
	return set
}

// Clear remove all elements with saving capacity (previous size)
func (set *Set[T]) Clear() *Set[T] {
	set.items = make([]T, 0, set.Size())
	return set
}

// Empty check is empty or not
func (set *Set[T]) Empty() bool {
	return set.Size() == 0
}

// Size returns size of collection
func (set *Set[T]) Size() int {
	return len(set.items)
}

// Get returns element on index
func (set *Set[T]) Get(index int) T {
	size := set.Size()

	if index > size || set.Empty() {
		emptyItem := new(T)
		return *emptyItem
	}

	return set.items[index]
}

// Sort implementation of sort package
func (set *Set[T]) Sort(compareFunc func(i, j int) bool) *Set[T] {
	sort.Slice(set.items, compareFunc)
	return set
}

// CopyTo copy collection to given
func (set *Set[T]) CopyTo(copyArray *Set[T]) *Set[T] {
	copyArray = NewWithList(set.items...)
	return set
}

// Reverse make reverse collection
func (set *Set[T]) Reverse() *Set[T] {
	for i, j := 0, len(set.items)-1; i < j; i, j = i+1, j-1 {
		set.items[i], set.items[j] = set.items[j], set.items[i]
	}

	return set
}

// Next moves iterator
func (set *Set[T]) Next() bool {
	iterator := set.iterator
	size := set.Size()

	if size == 0 || iterator == size {
		return false
	}

	set.iterator++
	return true
}

// Value returns value on iterator
func (set *Set[T]) Value() T {
	size := set.Size()
	if size == 0 {
		return *new(T)
	}

	return set.items[set.iterator-1]
}

// Reset set iterator to the start
func (set *Set[T]) Reset() *Set[T] {
	set.iterator = 0
	return set
}

// Clone create new collection with the same data
func (set *Set[T]) Clone() *Set[T] {
	return NewWithList[T](set.items...)
}

// CloneRange clone collection but with range
func (set *Set[T]) CloneRange(from, to int) *Set[T] {
	if from < 0 {
		from = 0
	}

	if set.indexOut(to) {
		to = set.Size()
	}

	return NewWithList[T](set.SliceRange(from, to)...)
}

// IndexOf search element and returns index
func (set *Set[T]) IndexOf(item T) int {
	return set.find(item)
}

// Contains true if collection owns element
func (set *Set[T]) Contains(item T) bool {
	return set.find(item) != -1
}

// ContainsFunc check for exist with given func
func (set *Set[T]) ContainsFunc(conditionFunc func(iterator T) bool) bool {
	for i := 0; i < set.Size(); i++ {
		if conditionFunc(set.items[i]) {
			return true
		}
	}

	return false
}

// Slice returns slice of collection
func (set *Set[T]) Slice() []T {
	return set.items
}

// SliceRange returns slice of collection with range
func (set *Set[T]) SliceRange(from, to int) []T {
	if from < 0 {
		from = 0
	}

	if set.indexOut(to) {
		to = set.Size()
	}

	return set.items[from:to]
}

// Select parse collection to another one with similar data type
func (set *Set[T]) Select(selectFunc func(sourceItem T) any) *Set[any] {
	newCollection := NewWithSize[any](set.Size())
	for _, item := range set.items {
		newCollection.Push(selectFunc(item))
	}
	return newCollection
}

// Sub returns new collection with range from this one
func (set *Set[T]) Sub(from, to int) *Set[T] {
	if from < 0 {
		from = 0
	}

	if set.indexOut(to) {
		to = set.Size()
	}

	return NewWithList(set.SliceRange(from, to)...)
}

package set

import (
	"github.com/lowl11/lazy-collection/helper"
	"sort"
)

func (set *Set[T]) Push(value T) *Set[T] {
	if helper.In[T](set.items, value) {
		return set
	}

	set.items = append(set.items, value)
	return set
}

func (set *Set[T]) PushForward(value T) *Set[T] {
	if helper.In[T](set.items, value) {
		return set
	}

	set.items = append([]T{value}, set.items...)
	return set
}

func (set *Set[T]) PushList(items ...T) *Set[T] {
	for _, item := range items {
		if !helper.In[T](set.items, item) {
			set.items = append(set.items, item)
		}
	}

	return set
}

func (set *Set[T]) Pop() *Set[T] {
	if set.Empty() {
		return set
	}

	set.items = set.items[:set.Size()-1]
	return set
}

func (set *Set[T]) PopForward() *Set[T] {
	if set.Empty() {
		return set
	}

	set.items = set.items[1:]
	return set
}

func (set *Set[T]) Remove(index int) *Set[T] {
	if set.Empty() || set.indexOut(index) {
		return set
	}

	set.items = append(set.items[:index], set.items[index+1:]...)
	return set
}

func (set *Set[T]) Empty() bool {
	return set.Size() == 0
}

func (set *Set[T]) Size() int {
	return len(set.items)
}

func (set *Set[T]) Get(index int) T {
	size := set.Size()

	if index > size || set.Empty() {
		emptyItem := new(T)
		return *emptyItem
	}

	return set.items[index]
}

func (set *Set[T]) Sort(compareFunc func(i, j int) bool) *Set[T] {
	sort.Slice(set.items, compareFunc)
	return set
}

func (set *Set[T]) Next() bool {
	iterator := set.iterator
	size := set.Size()

	if size == 0 || iterator == size {
		return false
	}

	set.iterator++
	return true
}

func (set *Set[T]) Value() T {
	size := set.Size()
	if size == 0 {
		return *new(T)
	}

	return set.items[set.iterator-1]
}

func (set *Set[T]) Reset() *Set[T] {
	set.iterator = 0
	return set
}

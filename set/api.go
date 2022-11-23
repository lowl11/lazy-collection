package set

import (
	"github.com/lowl11/lazy-collection/helper"
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

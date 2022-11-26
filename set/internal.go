package set

import "github.com/lowl11/lazy-collection/helper"

func (set *Set[T]) indexOut(index int) bool {
	return index >= set.Size()
}

func (set *Set[T]) find(value T) int {
	return helper.IndexOf[T](set.items, value)
}

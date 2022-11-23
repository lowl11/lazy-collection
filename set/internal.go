package set

func (set *Set[T]) indexOut(index int) bool {
	return index >= set.Size()
}

package set

func (set *Set[T]) indexOut(index int) bool {
	return index >= set.Size()
}

func (set *Set[T]) find(value T) int {
	for index, item := range set.items {
		if any(item) == any(value) {
			return index
		}
	}

	return -1
}

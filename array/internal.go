package array

func (array *Array[T]) indexOut(index int) bool {
	return index >= array.Size()
}

func (array *Array[T]) find(value T) int {
	for index, item := range array.items {
		if any(item) == any(value) {
			return index
		}
	}

	return -1
}

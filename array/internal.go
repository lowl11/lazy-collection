package array

func (array *Array[T]) indexOut(index int) bool {
	return index >= array.Size()
}

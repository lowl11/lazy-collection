package type_list

func (typeList *TypeList[T, X]) indexOut(index int) bool {
	return index >= typeList.Size()
}

func (typeList *TypeList[T, X]) find(value T) int {
	for index, item := range typeList.items {
		if any(item) == any(value) {
			return index
		}
	}

	return -1
}

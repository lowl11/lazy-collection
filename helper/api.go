package helper

func In[T any](collection []T, value T) bool {
	for _, item := range collection {
		if any(item) == any(value) {
			return true
		}
	}

	return false
}

func IndexOf[T any](collection []T, value T) int {
	for index, item := range collection {
		if any(item) == any(value) {
			return index
		}
	}

	return -1
}

func Fill[T any](collection []T, fill []T) {
	for index, item := range fill {
		collection[index] = item
	}
}

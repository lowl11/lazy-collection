package helper

type IComparable interface {
	int | string | float32 | float64 | rune
}

func In[T any](collection []T, value T) bool {
	for _, item := range collection {
		if any(item) == any(value) {
			return true
		}
	}

	return false
}

func Fill[T any](collection []T, fill []T) {
	for index, item := range fill {
		collection[index] = item
	}
}

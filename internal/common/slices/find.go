package slices

func Find[T any](values []T, predicate func(T) bool) (T, bool) {
	for _, el := range values {
		if predicate(el) {
			return el, true
		}
	}
	var zeroVal T
	return zeroVal, false
}

func Filter[T any, Slice ~[]T](collection Slice, predicate func(item T, index int) bool) Slice {
	result := make(Slice, 0, len(collection))

	for i := range collection {
		if predicate(collection[i], i) {
			result = append(result, collection[i])
		}
	}

	return result
}

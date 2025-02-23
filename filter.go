package collections

// Filter returns a new collection containing only the elements that satisfy the predicate function.
func Filter[T any](collection []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

package collections

// Count counts the number of elements in the collection that match the predicate.
func Count[T any](collection []T, predicate func(T) bool) int {
	count := 0
	for _, item := range collection {
		if predicate(item) {
			count++
		}
	}
	return count
}

package collections

// Any checks if any element in the collection matches the predicate.
func Any[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

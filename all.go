package collections

// All checks if all elements in the collection match the predicate.
func All[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if !predicate(item) {
			return false
		}
	}
	return true
}

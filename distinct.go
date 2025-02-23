package collections

// Distinct removes duplicate elements from a collection.
func Distinct[T comparable](collection []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if _, found := seen[item]; !found {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

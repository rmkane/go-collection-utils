package collections

// GroupBy groups elements in a collection by a specified key.
func GroupBy[T any, K comparable](collection []T, keyFunc func(T) K) map[K][]T {
	return Reduce(collection, func(acc map[K][]T, item T) map[K][]T {
		key := keyFunc(item)
		acc[key] = append(acc[key], item)
		return acc
	}, make(map[K][]T))
}

package collections

// FlatMap transforms each element in a collection and flattens the result.
func FlatMap[T any, R any](collection []T, transform func(T) []R) []R {
	return Reduce(collection, func(acc []R, item T) []R {
		return append(acc, transform(item)...)
	}, []R{})
}

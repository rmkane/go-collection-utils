package collections

import "golang.org/x/exp/constraints"

// Sum calculates the sum of elements in a collection.
func Sum[T constraints.Ordered](collection []T) T {
	var sum T
	return Reduce(collection, addFn, sum)
}

func addFn[T constraints.Ordered](a, b T) T {
	return a + b
}

package collections

import "golang.org/x/exp/constraints"

// Product calculates the product of elements in a collection.
func Product[T constraints.Integer | constraints.Float](collection []T) T {
	if len(collection) == 0 {
		var zero T
		return zero
	}
	return Reduce(collection, multiplyFn, 1.0)
}

func multiplyFn[T constraints.Integer | constraints.Float](a, b T) T {
	return a * b
}

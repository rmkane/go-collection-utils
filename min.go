package collections

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](collection []T) (T, bool) {
	if len(collection) == 0 {
		var zero T
		return zero, false
	}
	maxVal := Reduce(collection, minFn, collection[0])
	return maxVal, true
}

func minFn[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

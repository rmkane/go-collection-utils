package collections

import "golang.org/x/exp/constraints"

// Range calculates the range (difference between the maximum and minimum values) of elements in a collection.
func Range[T constraints.Integer | constraints.Float](collection []T) (T, bool) {
	if len(collection) == 0 {
		var zero T
		return zero, false
	}
	minVal, maxVal := collection[0], collection[0]
	for _, item := range collection {
		if item < minVal {
			minVal = item
		}
		if item > maxVal {
			maxVal = item
		}
	}
	return maxVal - minVal, true
}

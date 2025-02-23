package collections

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// Median calculates the median of elements in a collection.
func Median[T constraints.Integer | constraints.Float](collection []T) (T, bool) {
	if len(collection) == 0 {
		var zero T
		return zero, false
	}
	sorted := make([]T, len(collection))
	copy(sorted, collection)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2, true
	}
	return sorted[mid], true
}

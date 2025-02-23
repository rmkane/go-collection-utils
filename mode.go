package collections

import (
	"golang.org/x/exp/constraints"
)

// Mode calculates the mode (most frequent element) of elements in a collection.
func Mode[T constraints.Ordered](collection []T) (T, bool) {
	if len(collection) == 0 {
		var zero T
		return zero, false
	}

	frequency := Reduce(collection, func(freqMap map[T]int, item T) map[T]int {
		freqMap[item]++
		return freqMap
	}, make(map[T]int))

	var mode T
	maxCount := 0
	for item, count := range frequency {
		if count > maxCount {
			maxCount = count
			mode = item
		}
	}
	return mode, true
}

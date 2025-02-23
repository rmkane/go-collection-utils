package collections

import (
	"math"

	"golang.org/x/exp/constraints"
)

// StandardDeviation calculates the standard deviation of elements in a collection.
func StandardDeviation[T constraints.Integer | constraints.Float](collection []T) (float64, bool) {
	variance, ok := Variance(collection)
	if !ok {
		return 0, false
	}
	return math.Sqrt(variance), true
}

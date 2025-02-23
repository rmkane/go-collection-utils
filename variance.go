package collections

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Variance calculates the variance of elements in a collection.
func Variance[T constraints.Integer | constraints.Float](collection []T) (float64, bool) {
	if len(collection) == 0 {
		return 0, false
	}

	mean, _ := Average(collection)
	sum := Reduce(collection, func(acc float64, item T) float64 {
		return acc + math.Pow(float64(item)-mean, 2)
	}, 0.0)

	return sum / float64(len(collection)), true
}

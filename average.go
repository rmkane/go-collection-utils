package collections

import "golang.org/x/exp/constraints"

// Average calculates the average of elements in a collection and returns it as a float64.
func Average[T constraints.Integer | constraints.Float](collection []T) (float64, bool) {
	if len(collection) == 0 {
		return 0, false
	}
	sum := Reduce(collection, func(total float64, item T) float64 {
		return total + float64(item)
	}, 0.0)
	return sum / float64(len(collection)), true
}

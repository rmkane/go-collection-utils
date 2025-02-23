package collections

// Reduce applies a reducer function to each element in the collection, starting with an initial value, and returns the accumulated result.
func Reduce[T any, U any](collection []T, reducer func(U, T) U, initial U) U {
	accumulator := initial
	for _, item := range collection {
		accumulator = reducer(accumulator, item)
	}
	return accumulator
}

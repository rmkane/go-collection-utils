package collections

// Map applies a mapper function to each element in the collection and returns a new collection.
func Map[T any, U any](collection []T, mapper func(T) U) []U {
	result := make([]U, 0, len(collection))
	for _, item := range collection {
		result = append(result, mapper(item))
	}
	return result
}

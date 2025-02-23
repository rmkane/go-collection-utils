package collections

func EmptyCollection[T any]() []T {
	return make([]T, 0)
}

var EmptyIntCollection = EmptyCollection[int]()

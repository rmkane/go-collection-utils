package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_ContainsEvenNumbers(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}
	result := Filter(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.Equal(t, expected, result)
}

func TestFilter_NoEvenNumbers(t *testing.T) {
	collection := []int{1, 3, 5, 7, 9}
	expected := EmptyIntCollection
	result := Filter(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.Equal(t, expected, result)
}

func TestFilter_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	expected := EmptyIntCollection
	result := Filter(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.Equal(t, expected, result)
}

func TestFilter_AllEvenNumbers(t *testing.T) {
	collection := []int{2, 4, 6, 8, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := Filter(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.Equal(t, expected, result)
}

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount_ContainsEvenNumbers(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5, 6}
	result := Count(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.Equal(t, 3, result)
}

func TestCount_NoEvenNumbers(t *testing.T) {
	collection := []int{1, 3, 5, 7, 9}
	result := Count(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.Equal(t, 0, result)
}

func TestCount_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	result := Count(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.Equal(t, 0, result)
}

func TestCount_AllEvenNumbers(t *testing.T) {
	collection := []int{2, 4, 6, 8, 10}
	result := Count(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.Equal(t, 5, result)
}

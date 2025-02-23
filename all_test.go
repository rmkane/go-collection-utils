package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll_AllEvenNumbers(t *testing.T) {
	collection := []int{2, 4, 6, 8, 10}
	result := All(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.True(t, result)
}

func TestAll_ContainsOddNumber(t *testing.T) {
	collection := []int{2, 4, 6, 7, 10}
	result := All(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.False(t, result)
}

func TestAll_EmptyCollection(t *testing.T) {
	collection := EmptyCollection
	result := All(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.True(t, result) // An empty collection should return true
}

func TestAll_AllOddNumbers(t *testing.T) {
	collection := []int{1, 3, 5, 7, 9}
	result := All(collection, func(x int) bool {
		return x%2 != 0
	})
	assert.True(t, result)
}

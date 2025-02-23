package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains_NoEvenNumbers(t *testing.T) {
	collection := []int{1, 3, 5, 7, 9}
	result := Contains(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.False(t, result)
}

func TestContains_ContainsEvenNumber(t *testing.T) {
	collection := []int{1, 3, 5, 6, 9}
	result := Contains(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.True(t, result)
}

func TestContains_EmptyCollection(t *testing.T) {
	collection := EmptyCollection
	result := Contains(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.False(t, result) // An empty collection should return false
}

func TestContains_AllEvenNumbers(t *testing.T) {
	collection := []int{2, 4, 6, 8, 10}
	result := Contains(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.True(t, result)
}

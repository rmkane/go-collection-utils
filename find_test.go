package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind_NoEvenNumbers(t *testing.T) {
	collection := []int{1, 3, 5, 7, 9}
	result, found := Find(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.False(t, found)
	assert.Equal(t, 0, result)
}

func TestFind_ContainsEvenNumber(t *testing.T) {
	collection := []int{1, 3, 5, 6, 9}
	result, found := Find(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.True(t, found)
	assert.Equal(t, 6, result)
}

func TestFind_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	result, found := Find(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.False(t, found)
	assert.Equal(t, 0, result)
}

func TestFind_AllEvenNumbers(t *testing.T) {
	collection := []int{2, 4, 6, 8, 10}
	result, found := Find(collection, func(x int) bool {
		return x%2 == 0
	})
	assert.True(t, found)
	assert.Equal(t, 2, result)
}

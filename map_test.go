package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapCollection(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}
	result := Map(collection, func(x int) int {
		return x * 2
	})
	assert.Equal(t, expected, result)
}

func TestMap_EmptyCollection(t *testing.T) {
	collection := EmptyCollection
	expected := EmptyCollection
	result := Map(collection, func(x int) int {
		return x * 2
	})
	assert.Equal(t, expected, result)
}

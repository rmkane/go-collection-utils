package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceCollection(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	expected := 15
	result := Reduce(collection, func(acc, x int) int {
		return acc + x
	}, 0)
	assert.Equal(t, expected, result)
}

func TestReduce_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	expected := 0
	result := Reduce(collection, func(acc, x int) int {
		return acc + x
	}, 0)
	assert.Equal(t, expected, result)
}

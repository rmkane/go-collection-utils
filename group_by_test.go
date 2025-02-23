package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy_IntCollection(t *testing.T) {
	collection := []int{1, 2, 2, 3, 4, 4, 4, 5}
	grouped := GroupBy(collection, func(n int) int { return n % 2 })
	assert.Equal(t, map[int][]int{
		0: {2, 2, 4, 4, 4},
		1: {1, 3, 5},
	}, grouped)
}

func TestGroupBy_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	grouped := GroupBy(collection, func(n int) int { return n % 2 })
	assert.Equal(t, map[int][]int{}, grouped)
}

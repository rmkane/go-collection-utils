package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	collection := []int{1, 2, 3}
	transform := func(x int) []int {
		return []int{x, x * 2}
	}
	result := FlatMap(collection, transform)
	expected := []int{1, 2, 2, 4, 3, 6}
	assert.Equal(t, expected, result)
}

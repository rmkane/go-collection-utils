package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange_IntCollection(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	rng, ok := Range(collection)
	assert.True(t, ok)
	assert.Equal(t, 4, rng)
}

func TestRange_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	rng, ok := Range(collection)
	assert.False(t, ok)
	assert.Equal(t, 0, rng)
}

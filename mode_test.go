package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMode_IntCollection(t *testing.T) {
	collection := []int{1, 2, 2, 3, 4, 4, 4, 5}
	mode, ok := Mode(collection)
	assert.True(t, ok)
	assert.Equal(t, 4, mode)
}

func TestMode_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	mode, ok := Mode(collection)
	assert.False(t, ok)
	assert.Equal(t, 0, mode)
}

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMin_IntCollection(t *testing.T) {
	collection := []int{5, 3, 9, 1, 4}
	minVal, ok := Min(collection)
	assert.True(t, ok)
	assert.Equal(t, 1, minVal)
}

func TestMin_FloatCollection(t *testing.T) {
	collection := []float64{5.5, 3.3, 9.9, 1.1, 4.4}
	minVal, ok := Min(collection)
	assert.True(t, ok)
	assert.Equal(t, 1.1, minVal)
}

func TestMin_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	minVal, ok := Min(collection)
	assert.False(t, ok)
	assert.Equal(t, 0, minVal)
}

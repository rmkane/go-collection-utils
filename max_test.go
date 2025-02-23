package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax_IntCollection(t *testing.T) {
	collection := []int{5, 3, 9, 1, 4}
	maxVal, ok := Max(collection)
	assert.True(t, ok)
	assert.Equal(t, 9, maxVal)
}

func TestMax_FloatCollection(t *testing.T) {
	collection := []float64{5.5, 3.3, 9.9, 1.1, 4.4}
	maxVal, ok := Max(collection)
	assert.True(t, ok)
	assert.Equal(t, 9.9, maxVal)
}

func TestMax_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	maxVal, ok := Max(collection)
	assert.False(t, ok)
	assert.Equal(t, 0, maxVal)
}

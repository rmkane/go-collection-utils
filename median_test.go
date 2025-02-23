package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedian_IntCollection(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	median, ok := Median(collection)
	assert.True(t, ok)
	assert.Equal(t, 3, median)
}

func TestMedian_FloatCollection(t *testing.T) {
	collection := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	median, ok := Median(collection)
	assert.True(t, ok)
	assert.Equal(t, 3.3, median)
}

func TestMedian_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	median, ok := Median(collection)
	assert.False(t, ok)
	assert.Equal(t, 0, median)
}

package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage_IntCollection(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	avg, ok := Average(collection)
	assert.True(t, ok)
	assert.Equal(t, 3.0, avg)
}

func TestAverage_FloatCollection(t *testing.T) {
	collection := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	avg, ok := Average(collection)
	assert.True(t, ok)
	assert.Equal(t, 3.3, avg)
}

func TestAverage_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	avg, ok := Average(collection)
	assert.False(t, ok)
	assert.Equal(t, 0.0, avg)
}

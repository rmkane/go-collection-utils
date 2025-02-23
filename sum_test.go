package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum_IntCollection(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	sum := Sum(collection)
	assert.Equal(t, 15, sum)
}

func TestSum_FloatCollection(t *testing.T) {
	collection := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	sum := Sum(collection)
	assert.Equal(t, 16.5, sum)
}

func TestSum_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	sum := Sum(collection)
	assert.Equal(t, 0, sum)
}

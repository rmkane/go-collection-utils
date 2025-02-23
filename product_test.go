package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct_IntCollection(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5}
	product := Product(collection)
	assert.Equal(t, 120, product)
}

func TestProduct_FloatCollection(t *testing.T) {
	collection := []float64{1.1, 2.2, 3.3}
	product := Product(collection)
	assert.InDelta(t, 7.986, product, 0.001)
}

func TestProduct_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	product := Product(collection)
	assert.Equal(t, 0, product)
}

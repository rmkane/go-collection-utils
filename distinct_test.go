package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistinct_IntCollection(t *testing.T) {
	collection := []int{1, 2, 2, 3, 4, 4, 4, 5}
	distinct := Distinct(collection)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, distinct)
}

func TestDistinct_EmptyCollection(t *testing.T) {
	collection := EmptyIntCollection
	distinct := Distinct(collection)
	assert.Equal(t, []int{}, distinct)
}

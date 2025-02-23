package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariance(t *testing.T) {
	tests := []struct {
		name       string
		collection []int
		expected   float64
		ok         bool
	}{
		{
			name:       "Non-empty collection",
			collection: []int{1, 2, 3, 4, 5},
			expected:   2.0,
			ok:         true,
		},
		{
			name:       "Single element collection",
			collection: []int{42},
			expected:   0.0,
			ok:         true,
		},
		{
			name:       "Empty collection",
			collection: []int{},
			expected:   0.0,
			ok:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := Variance(tt.collection)
			assert.Equal(t, tt.ok, ok)
			assert.InDelta(t, tt.expected, result, 1e-9)
		})
	}
}

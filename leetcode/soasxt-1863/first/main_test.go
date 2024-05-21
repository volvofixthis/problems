package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   []int
		o2   int
	}{
		{"first", []int{1, 3}, 6},
		{"second", []int{5, 1, 6}, 28},
		// {"third", []int{3, 4, 5, 6, 7, 8}, 480},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o2, subsetXORSum(test.i1), "should be equal")
		})
	}
}

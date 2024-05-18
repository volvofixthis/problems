package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   []int
		o1   []int
	}{
		{"one", []int{0}, []int{0}},
		{"two", []int{3, 2}, []int{2, 3}},
		{"tree", []int{2, 4, 3, 1}, []int{2, 4, 1, 3}},
		{"tree", []int{2, 4, 3, 5, 1}, []int{2, 4, 1, 3, 5}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, sortArrayByParity(test.i1), "should be equal")
		})
	}
}

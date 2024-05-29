package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 string
		o1 int
	}{
		{"1101", 6},
		{"10", 1},
		{"1", 0},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, numSteps(test.i1), "should be equal")
		})
	}
}

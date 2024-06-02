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
		{"EEEEEEE", 7},
		{"ELELEEL", 2},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, minimumChairs(test.i1), "should be equal")
		})
	}
}

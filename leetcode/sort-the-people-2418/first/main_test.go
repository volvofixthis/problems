package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   []string
		i2   []int
		o1   []string
	}{
		{"first", []string{"Mary", "John", "Emma"}, []int{180, 165, 170}, []string{"Mary", "Emma", "John"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, sortPeople(test.i1, test.i2), "should be equal")
		})
	}
}

package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   string
		o1   [][]string
	}{
		{"one", "aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"one", "a", [][]string{{"a"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, partition(test.i1), "should be equal")
		})
	}
}

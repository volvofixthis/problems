package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	i1 string
	i2 string
	o1 int
}{
	{"coaching", "coding", 4},
	{"abcde", "a", 0},
	{"z", "abcde", 5},
}

func TestFunc(t *testing.T) {
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, appendCharacters(test.i1, test.i2), "should be equal")
		})
	}
}

package first

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	tests := []struct {
		i1 []string
		i2 string
		o1 string
	}{
		{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery", "the cat was rat by the bat"},
		{[]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs", "a a b c"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, replaceWords(test.i1, test.i2), "should be equal")
		})
	}
}

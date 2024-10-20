package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripSpaces(t *testing.T) {
	tests := []struct {
		i1 []rune
		o1 []rune
	}{
		{[]rune{rune(' ')}, []rune{rune(' ')}},
		{[]rune{rune('a'), rune(' '), rune(' '), rune('b')}, []rune{rune('a'), rune(' '), rune('b')}},
		{[]rune{rune('a')}, []rune{rune('a')}},
		{[]rune{rune(' '), rune('a')}, []rune{rune(' '), rune('a')}},
		{[]rune{rune('a'), rune(' ')}, []rune{rune('a'), rune(' ')}},
		{[]rune{rune('a'), rune(' '), rune(' ')}, []rune{rune('a'), rune(' ')}},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			assert.Equal(t, test.o1, stripSpaces(test.i1), "must be equal")
		})
	}
}

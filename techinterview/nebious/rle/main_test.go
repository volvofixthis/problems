package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRle(t *testing.T) {
	tests := []struct {
		i1 string
		o1 string
		o2 error
	}{
		{"AAAABBB", "A4B3", nil},
		{"A", "A", nil},
		{"AA", "A2", nil},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1),
			func(t *testing.T) {
				a1, err := rle(test.i1)
				assert.Equal(t, test.o1, a1, "Must be equal")
				assert.Equal(t, test.o2, err, "Must be equal")
			},
		)
	}
}

func TestRle2(t *testing.T) {
	tests := []struct {
		i1 string
		o1 string
		o2 error
	}{
		{"AAAABBB", "A4B3", nil},
		{"A", "A", nil},
		{"AA", "A2", nil},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1),
			func(t *testing.T) {
				a1, err := rle2(test.i1)
				assert.Equal(t, test.o1, a1, "Must be equal")
				assert.Equal(t, test.o2, err, "Must be equal")
			},
		)
	}
}

func TestRle3(t *testing.T) {
	tests := []struct {
		i1 string
		o1 string
		o2 error
	}{
		{"AAAABBB", "A4B3", nil},
		{"A", "A", nil},
		{"AA", "A2", nil},
		{"1", "", ErrWrongSymbol},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1),
			func(t *testing.T) {
				a1, err := rle3(test.i1)
				assert.Equal(t, test.o1, a1, "Must be equal")
				assert.Equal(t, test.o2, err, "Must be equal")
			},
		)
	}
}

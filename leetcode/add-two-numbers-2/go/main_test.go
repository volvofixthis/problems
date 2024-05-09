package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCovertors(t *testing.T) {
	tests := []struct {
		name  string
		input int
	}{
		{"first", 342},
		{"second", 465},
		{"third", 9999},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, listToNumber(numberToList(test.input)), test.input, "should be equal")
		})
	}
}

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		l1   int
		l2   int
		l3   int
	}{
		{"first", 342, 465, 807},
		{"second", 0, 0, 0},
		{"third", 999, 99, 1098},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, listToNumber(addTwoNumbers(numberToList(test.l1), numberToList(test.l2))), test.l3, "should be equal")
		})
	}
}

func TestBigFunc(t *testing.T) {
	tests := []struct {
		name string
		l1   string
		l2   string
		l3   string
	}{
		{"first", "2,4,3", "5,6,4", "7,0,8"},
		{
			"second",
			"5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1",
			"9,4,5",
			"4,5,5,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, strToList(test.l3).Val, addTwoNumbers(strToList(test.l1), strToList(test.l2)).Val, "should be equal")
		})
	}
}

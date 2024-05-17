package first

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkFunc(b *testing.B) {
	tests := []struct {
		name string
		i1   string
		i2   int
		o1   string
	}{
		{"one", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"two", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"three", "A", 1, "A"},
		{"four", "AB", 2, "AB"},
		{"five", "ABC", 2, "ACB"},
		{"six", "1234567", 3, "1524637"},
	}

	for _, test := range tests {
		convert(test.i1, test.i2)
	}
}

func TestFunc(t *testing.T) {
	tests := []struct {
		name string
		i1   string
		i2   int
		o1   string
	}{
		{"one", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"two", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"three", "A", 1, "A"},
		{"four", "AB", 2, "AB"},
		{"five", "ABC", 2, "ACB"},
		{"six", "1234567", 3, "1524637"},
		{"seven", "AB", 1, "AB"},
		{"eight", "ABCD", 2, "ACBD"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.o1, convert(test.i1, test.i2), "should be equal")
		})
	}
}

package main

import "testing"

// OneEditApart("cat", "dog") -> false
// OneEditApart("cat", "cats") -> true
// OneEditApart("cat", "cut") -> true
// OneEditApart("cat", "cast") -> true
// OneEditApart("cat", "at") -> true
// OneEditApart("cat", "acts") -> false
// OneEditApart("cat", "act") -> false

func TestOneEditApart(t *testing.T) {
	tests := []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"cat", "dog", false},
		{"cat", "cats", true},
		{"cat", "cut", true},
		{"cat", "cast", true},
		{"cat", "at", true},
		{"cat", "acts", false},
		{"cat", "act", false},
	}

	for _, test := range tests {
		if result := OneEditApart(test.str1, test.str2); result != test.expected {
			t.Errorf("OneEditApart(%q, %q) = %v; want %v", test.str1, test.str2, result, test.expected)
		}
	}
}

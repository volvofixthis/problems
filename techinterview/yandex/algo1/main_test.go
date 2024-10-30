package main

// 1. normalize('/foo/bar//baz/asdf/quux/..');         // '/foo/bar/baz/asdf'
// 2. normalize('./config/../etc');                    // 'etc'
// 3. normalize('/////documents/root/.././../etc');    // '/etc'
// 4. normalize('a/../../b');                          // '../b'
// 5. normalize('/a/../../b');                         // '/b'
import (
	"testing"
)

func TestNormalize(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"/foo/bar//baz/asdf/quux/..", "/foo/bar/baz/asdf"},
		{"./config/../etc", "etc"},
		{"/////documents/root/.././../etc", "/etc"},
		{"a/../../b", "../b"},
		{"/a/../../b", "/b"},
	}

	for _, c := range cases {
		result := normalize(c.input)
		if result != c.expected {
			t.Errorf("normalize(%q) == %q, expected %q", c.input, result, c.expected)
		}
	}
}

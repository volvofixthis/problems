package main

import (
	"reflect"
	"testing"
)

// Example 1:
//
// Input: folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
// Output: ["/a","/c/d","/c/f"]
// Explanation: Folders "/a/b" is a subfolder of "/a" and "/c/d/e" is inside of folder "/c/d" in our filesystem.
//
// Example 2:
//
// Input: folder = ["/a","/a/b/c","/a/b/d"]
// Output: ["/a"]
// Explanation: Folders "/a/b/c" and "/a/b/d" will be removed because they are subfolders of "/a".
//
// Example 3:
//
// Input: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
// Output: ["/a/b/c","/a/b/ca","/a/b/d"]

func TestRemoveSubfolders(t *testing.T) {
	tests := []struct{ i1, i2 []string }{
		{[]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}, []string{"/a", "/c/d", "/c/f"}},
		{[]string{"/a", "/a/b/c", "/a/b/d"}, []string{"/a"}},
		{[]string{"/a/b/c", "/a/b/ca", "/a/b/d"}, []string{"/a/b/c", "/a/b/ca", "/a/b/d"}},
	}

	for _, test := range tests {
		if result := removeSubfolders(test.i1); !reflect.DeepEqual(result, test.i2) {
			t.Errorf("removeSubfolders(%v) = %v; expected %v", test.i1, result, test.i2)
		}
	}
}

package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Example 1:
//
// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
//
// Output: ["facebook","google","leetcode"]
//
// Example 2:
//
// Input: words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["lc","eo"]
//
// Output: ["leetcode"]
//
// Example 3:
//
// Input: words1 = ["acaac","cccbb","aacbb","caacc","bcbbb"], words2 = ["c","cc","b"]
//
// Output: ["cccbb"]

func TestWordSubsets(t *testing.T) {
	tests := []struct {
		i1 []string
		i2 []string
		o1 []string
	}{
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}, []string{"facebook", "google", "leetcode"}},
		{[]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"lc", "eo"}, []string{"leetcode"}},
		{[]string{"acaac", "cccbb", "aacbb", "caacc", "bcbbb"}, []string{"c", "cc", "b"}, []string{"cccbb"}},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			if v := wordSubsets(test.i1, test.i2); !reflect.DeepEqual(v, test.o1) {
				t.Errorf("Result %s isn't equal to output %s", v, test.o1)
			}
		})
	}
}

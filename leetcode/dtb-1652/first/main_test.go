package main

import (
	"fmt"
	"reflect"
	"testing"
)

// You have a bomb to defuse, and your time is running out! Your informer will provide you with a circular array code of length of n and a key k.
//
// To decrypt the code, you must replace every number. All the numbers are replaced simultaneously.
//
//     If k > 0, replace the ith number with the sum of the next k numbers.
//     If k < 0, replace the ith number with the sum of the previous k numbers.
//     If k == 0, replace the ith number with 0.
//
// As code is circular, the next element of code[n-1] is code[0], and the previous element of code[0] is code[n-1].
//
// Given the circular array code and an integer key k, return the decrypted code to defuse the bomb!

// Example 1:
//
// Input: code = [5,7,1,4], k = 3
// Output: [12,10,16,13]
// Explanation: Each number is replaced by the sum of the next 3 numbers. The decrypted code is [7+1+4, 1+4+5, 4+5+7, 5+7+1]. Notice that the numbers wrap around.
//
// Example 2:
//
// Input: code = [1,2,3,4], k = 0
// Output: [0,0,0,0]
// Explanation: When k is zero, the numbers are replaced by 0.
//
// Example 3:
//
// Input: code = [2,4,9,3], k = -2
// Output: [12,5,6,13]
// Explanation: The decrypted code is [3+9, 2+3, 4+2, 9+4]. Notice that the numbers wrap around again. If k is negative, the sum is of the previous numbers.
//

func TestDecrypt(t *testing.T) {
	tests := []struct {
		i1 []int
		i2 int
		o1 []int
	}{
		{[]int{5, 7, 1, 4}, 3, []int{12, 10, 16, 13}},
		{[]int{1, 2, 3, 4}, 0, []int{0, 0, 0, 0}},
		{[]int{2, 4, 9, 3}, -2, []int{12, 5, 6, 13}},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i+1), func(t *testing.T) {
			if v := decrypt(test.i1, test.i2); !reflect.DeepEqual(test.o1, v) {
				t.Errorf("Result %v isn't equal to %v", v, test.o1)
			}
		})
	}
}

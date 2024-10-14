package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 5, 6, 7, 8, 18, 99} // слайсы отсортированы
	var b = []int{2, 4, 9, 11}
	fmt.Println(mergeSorted(a, b))
}

func mergeSorted(a []int, b []int) []int {
	// напишите код который который смержит слайсы, a/b и с затратами памяти O(n), сложности O(n)
	ap := 0
	bp := 0
	result := make([]int, 0, len(a)+len(b))
	for {
		if ap < len(a) && (bp >= len(b) || a[ap] < b[bp]) {
			result = append(result, a[ap])
			ap++
		} else if bp < len(b) {
			result = append(result, b[bp])
			bp++
		} else {
			break
		}
	}
	return result
}

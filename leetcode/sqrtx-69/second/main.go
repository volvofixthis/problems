package main

func mySqrt(x int) int {
	left := 1
	right := x/2 + 1

	for left <= right {
		mid := left + (right-left)/2
		p := mid * mid
		if p < x {
			left = mid + 1
		} else if p > x {
			right = mid - 1
		} else {
			return mid
		}
	}
	return right
}

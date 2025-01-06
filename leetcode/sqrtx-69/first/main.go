package main

func mySqrt(x int) int {
	half := x/2 + 1
	target := 0
	for i := 0; i <= half; i++ {
		if i*i <= x {
			target = i
		} else {
			break
		}
	}
	return target
}

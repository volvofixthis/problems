package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	order := 1
	for x > order {
		order *= 10
	}
	if order > x {
		order = order / 10
	}
	result := 0
	cp := order
	c := x
	rp := 1
	for cp > 0 {
		digit := c / cp
		result += digit * rp
		c = c % cp
		rp = rp * 10
		cp = cp / 10
	}
	return result == x
}

package main

import "slices"

func AbsDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func findReqNeg(goods []int, curReq int) int {
	l := 0
	r := len(goods)
	for l < r {
		mid := l + (r-l)/2
		if curReq < goods[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	results := []int{}
	if l-1 >= 0 {
		results = append(results, AbsDiff(goods[l-1], curReq))
	}
	if l < len(goods) {
		results = append(results, AbsDiff(goods[l], curReq))
	}
	return slices.Min(results)
}

func calcNeg(goods []int, req []int) int {
	s := 0
	for _, v := range req {
		s += findReqNeg(goods, v)
	}
	return s
}

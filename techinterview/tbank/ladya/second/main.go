package main

import (
	"fmt"
)

var a = [][]int{
	{1, 4, 1},
	{17, 10, 1},
	{3, 15, 2},
}

func getMaxDamage(a [][]int) int {
	if len(a) == 0 {
		return 0
	}
	rowssum := make([]int, len(a))
	colssum := make([]int, len(a[0]))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			rowssum[i] += a[i][j]
			colssum[j] += a[i][j]
		}
	}
	mr := 0
	mc := 0
	for i := 0; i < len(rowssum); i++ {
		if rowssum[mr] < rowssum[i] {
			mr = i
		}
	}
	for i := 0; i < len(colssum); i++ {
		if colssum[mc] < colssum[i] {
			mc = i
		}
	}

	return rowssum[mr] + colssum[mc] - a[mr][mc]
}

func main() {
	fmt.Println(getMaxDamage(a))
}

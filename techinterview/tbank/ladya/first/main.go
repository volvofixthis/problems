package main

import "fmt"

// Дана "шахматная доска" m x n. Нужно найти клетку где ладья бьёт максимальное количество
// чисел по горизонтали и вертикали

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
	mx := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			rowssum[i] += a[i][j]
			colssum[j] += a[i][j]
		}
	}
	for i := 0; i < len(rowssum); i++ {
		for j := 0; j < len(colssum); j++ {
			sum := rowssum[i] + colssum[j] - a[i][j]
			if mx < sum {
				mx = sum
			}
		}
	}
	return mx
}

func main() {
	fmt.Println(getMaxDamage(a))
}

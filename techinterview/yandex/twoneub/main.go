package main

import "fmt"

/*
Даны две отсортированных по неубыванию последовательности целых чисел.
Необходимо вернуть все элементы из первой последовательности, которых нет во второй.
Пример:
FilterSortedList([1, 2, 3], [3, 4]) -> [1, 2]
FilterSortedList([1, 2, 3, 5], [2, 3, 7]) -> [1, 5]
FilterSortedList([3, 5], []) -> [3, 5]
[1, 1, 1], [1, 2] -> []
// Не инкрементить bp
*/

func FilterSortedList(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	ap := 0
	bp := 0
	for ap < len(a) && bp < len(b) {
		if a[ap] < b[bp] {
			result = append(result, a[ap])
			ap++
		} else if a[ap] > b[bp] {
			bp++
		} else {
			for ap < len(a) && a[ap] == b[bp] {
				ap++
			}
			bp++
		}
	}
	for ap < len(a) {
		result = append(result, a[ap])
		ap++
	}
	return result
}

func main() {
	fmt.Println(FilterSortedList([]int{1, 2, 3, 5}, []int{2, 3, 7}))
	fmt.Println(FilterSortedList([]int{1, 1, 1}, []int{1, 2}))
}

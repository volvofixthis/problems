package first

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := map[int]int{}
	for i, v := range arr2 {
		m[v] = i
	}
	arrR := []int{}
	arrEnd := []int{}
	for _, v := range arr1 {
		if _, ok := m[v]; ok {
			arrR = append(arrR, v)
		} else {
			arrEnd = append(arrEnd, v)
		}
	}
	sort.Slice(arrR, func(i, j int) bool {
		if m[arrR[i]] < m[arrR[j]] {
			return true
		}
		return false
	})
	sort.Ints(arrEnd)
	arrR = append(arrR, arrEnd...)
	return arrR
}

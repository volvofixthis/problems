package first

func singleNumber(nums []int) []int {
	m := map[int]int{}
	result := []int{}
	for _, v := range nums {
		m[v] += 1
		if m[v] == 2 {
			delete(m, v)
		}
	}
	for k := range m {
		result = append(result, k)
	}
	return result
}

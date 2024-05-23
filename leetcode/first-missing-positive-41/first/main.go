package first

func firstMissingPositive(nums []int) int {
	existing := map[int]struct{}{}
	m := 1
	for _, v := range nums {
		if v > m {
			m = v
		}
		existing[v] = struct{}{}
	}
	for i := 1; i <= m+1; i++ {
		if _, ok := existing[i]; !ok {
			return i
		}
	}
	return 0
}

package first

func checkSubarraySum(nums []int, k int) bool {
	prefixSum := 0
	m := map[int]int{0: -1}
	for i := 0; i < len(nums); i++ {
		prefixSum = (prefixSum + nums[i]) % k
		if v, ok := m[prefixSum]; ok {
			if i-v > 1 {
				return true
			}
		} else {
			m[prefixSum] = i
		}
	}
	return false
}

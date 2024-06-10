package first

func subarraysDivByK(nums []int, k int) int {
	result := 0
	prefixMod := 0
	m := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		prefixMod = (prefixMod + nums[i]%k + k) % k
		result += m[prefixMod]
		m[prefixMod] += 1
	}
	return result
}

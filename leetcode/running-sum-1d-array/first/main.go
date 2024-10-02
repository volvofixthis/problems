package first

func runningSum(nums []int) []int {
	prefixSum := 0
	result := []int{}
	for _, v := range nums {
		prefixSum += v
		result = append(result, prefixSum)
	}
	return result
}

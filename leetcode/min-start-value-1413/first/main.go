package first

func minStartValue(nums []int) int {
	prefixSum := 0
	minValue := 0
	for _, v := range nums {
		prefixSum += v
		if prefixSum < minValue {
			minValue = prefixSum
		}
	}
	startValue := 1 - minValue
	return startValue
}

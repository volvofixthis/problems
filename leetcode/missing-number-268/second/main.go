package first

func missingNumber(nums []int) int {
	missing := len(nums)
	for i, v := range nums {
		missing ^= i ^ v
	}
	return missing
}

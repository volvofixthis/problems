package first

func missingNumber(nums []int) int {
	n := len(nums) + 1
	current := 0
	for _, v := range nums {
		current += v
	}
	real := 0
	for i := 1; i < n; i++ {
		real += i
	}
	return real - current
}

package first

func subsetXORSum(nums []int) int {
	x := 0
	for _, v := range nums {
		x |= v
	}
	x = x << (len(nums) - 1)
	return x
}

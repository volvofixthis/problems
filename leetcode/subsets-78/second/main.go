package first

var result = [][]int{}

func backtrack(nums []int, start int, current []int) {
	if len(nums) == start {
		result = append(result, current)
	} else {
		backtrack(nums, start+1, current)
		backtrack(nums, start+1, append(current, nums[start]))
	}
}

func subsets(nums []int) [][]int {
	result = result[:0]
	for i := 0; i < 1<<len(nums); i++ {
		a := []int{}
		for j := 0; j < len(nums); j++ {
			if (i>>j)&1 == 1 {
				a = append(a, nums[j])
			}
		}
		result = append(result, a)
	}
	return result
}

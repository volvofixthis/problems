package first

var result = [][]int{}

func backtrack(nums []int, start int, current []int) {
	result = append(result, current)
	for i := start; i < len(nums); i++ {
		backtrack(nums, i+1, append(current, nums[i]))
	}
}

func subsets(nums []int) [][]int {
	result = [][]int{}
	backtrack(nums, 0, []int{})
	return result
}

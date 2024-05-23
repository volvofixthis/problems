package first

import "math"

func countBeautifulSubsets(nums []int, difference int, index int, mask int) int {
	if index == len(nums) {
		if mask > 0 {
			return 1
		} else {
			return 0
		}
	}

	// Flag to check if the current subset is beautiful
	isBeautiful := true

	// Check if the current number forms a beautiful pair with any previous
	// number in the subset
	for j := 0; j < index && isBeautiful; j++ {
		isBeautiful = ((1<<j)&mask) == 0 ||
			math.Abs(float64(nums[j]-nums[index])) != float64(difference)
	}

	// Recursively calculate beautiful subsets including and excluding the
	// current number
	skip := countBeautifulSubsets(nums, difference, index+1, mask)
	take := 0
	if isBeautiful {
		take = countBeautifulSubsets(nums, difference, index+1,
			mask+(1<<index))
	} else {
		take = 0
	}

	return skip + take
}

func beautifulSubsets(nums []int, k int) int {
	return countBeautifulSubsets(nums, k, 0, 0)
}

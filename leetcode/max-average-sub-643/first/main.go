package first

func findMaxAverage(nums []int, k int) float64 {
	window_start := 0
	window_sum := 0
	max_average := 0.0
	for i := 0; i < len(nums); i++ {
		window_sum += nums[i]
		if i-window_start >= k {
			window_sum -= nums[window_start]
			window_start += 1
		}
		cur_average := float64(window_sum) / float64(k)
		if cur_average > max_average {
			max_average = cur_average
		}
	}
	return max_average
}

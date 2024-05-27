package first

func specialArray(nums []int) int {
	n := len(nums)

	freq := make([]int, n+1)
	for _, v := range nums {
		freq[min(n, v)] += 1
	}
	s := 0
	for i := n; i > 0; i-- {
		s += freq[i]
		if i == s {
			return i
		}
	}

	return -1
}

package first

func sortColors(nums []int) {
	m := map[int]int{}
	for _, v := range nums {
		m[v] += 1
	}
	p := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < m[i]; j++ {
			nums[p] = i
			p++
		}
	}
}

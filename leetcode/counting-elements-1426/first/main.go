package first

func countElements(arr []int) int {
	st := map[int]int{}
	for _, v := range arr {
		st[v] += 1
	}
	found := 0
	for _, v := range arr {
		if _, ok := st[v+1]; ok {
			found += 1
		}
	}
	return found
}

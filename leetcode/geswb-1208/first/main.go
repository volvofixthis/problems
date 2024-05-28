package first

func absDiffInt(x, y byte) int {
	if x < y {
		return int(y - x)
	}
	return int(x - y)
}

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	start := 0
	curr_cost := 0
	max_len := 0
	for i := 0; i < n; i++ {
		curr_cost += absDiffInt(s[i], t[i])
		for curr_cost > maxCost {
			curr_cost -= absDiffInt(s[start], t[start])
			start += 1
		}
		max_len = max(max_len, i-start+1)
	}
	return max_len
}

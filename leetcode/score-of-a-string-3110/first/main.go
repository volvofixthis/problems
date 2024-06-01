package first

func absDiffInt(x, y byte) int {
	if x < y {
		return int(y - x)
	}
	return int(x - y)
}

func scoreOfString(s string) int {
	sum := 0
	for i := 1; i < len(s); i++ {
		sum += absDiffInt(s[i], s[i-1])
	}
	return sum
}

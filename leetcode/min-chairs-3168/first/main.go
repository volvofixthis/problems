package first

func minimumChairs(s string) int {
	chairs := 0
	maxChairs := 0
	for _, c := range s {
		if c == 'E' {
			chairs += 1
		} else {
			chairs -= 1
		}
		if chairs > maxChairs {
			maxChairs = chairs
		}
	}
	return maxChairs
}

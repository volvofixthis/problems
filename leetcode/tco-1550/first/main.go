package first

func threeConsecutiveOdds(arr []int) bool {
	oc := 0
	for _, v := range arr {
		if v%2 == 1 {
			oc += 1
			if oc >= 3 {
				return true
			}
		} else {
			oc = 0
		}
	}
	return false
}

package first

func countDays(days int, meetings [][]int) int {
	fdm := make([]int, days)
	for _, meeting := range meetings {
		for i := meeting[0]; i <= meeting[1]; i++ {
			fdm[i-1] = -1
		}
	}
	fd := 0
	for _, v := range fdm {
		if v >= 0 {
			fd += 1
		}
	}
	return fd
}

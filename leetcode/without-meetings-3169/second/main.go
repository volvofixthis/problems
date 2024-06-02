package first

import (
	"sort"
)

func countDays(days int, meetings [][]int) int {
	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	fd := 0
	b := 1
	for _, meeting := range meetings {
		if b < meeting[0] {
			fd += meeting[0] - b
		}
		if meeting[1] >= b {
			b = meeting[1] + 1
		}
	}
	if b <= days {
		fd += days - b + 1
	}
	return fd
}

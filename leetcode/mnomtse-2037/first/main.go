package first

import (
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	moves := 0
	sort.Ints(seats)
	sort.Ints(students)
	for i := 0; i < len(seats); i++ {
		diff := 0
		if students[i] < seats[i] {
			diff = seats[i] - students[i]
		} else {
			diff = students[i] - seats[i]
		}
		moves += diff
	}
	return moves
}

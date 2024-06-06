package first

import (
	"sort"
)

func isNStraightHand(hand []int, groupSize int) bool {
	nGroups := len(hand) / groupSize
	if len(hand)%groupSize != 0 {
		return false
	}
	sort.Ints(hand)
	m := map[int]int{}
	for _, v := range hand {
		m[v] += 1
	}

	ph := 0
	for i := 0; i < nGroups; i++ {
		n := hand[ph]
		for j := 0; j < groupSize; j++ {
			if v, ok := m[n]; !ok || v == 0 {
				return false
			}
			m[n] -= 1
			n += 1
		}
		for z := ph + 1; z < len(hand); z++ {
			if v, ok := m[hand[z]]; ok && v > 0 {
				ph = z
				break
			}
		}
	}
	return true
}

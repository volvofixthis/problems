package first

import "slices"

func sortPeople(names []string, heights []int) []string {
	type Pair struct {
		Name   string
		Height int
	}
	pairs := []Pair{}
	for k, v := range names {
		pairs = append(pairs, Pair{Name: v, Height: heights[k]})
	}
	slices.SortFunc(pairs, func(a, b Pair) int {
		if a.Height > b.Height {
			return -1
		}
		if a.Height < b.Height {
			return 1
		}
		return 0
	})
	result := []string{}
	for _, pair := range pairs {
		result = append(result, pair.Name)
	}
	return result
}

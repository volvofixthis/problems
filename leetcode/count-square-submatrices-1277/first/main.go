package main

func countSquares(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	prev := make([]int, cols+1)
	curr := make([]int, cols+1)
	r := 0
	for i := range rows {
		for j := 1; j < cols+1; j++ {
			if matrix[i][j-1] == 1 {
				curr[j] = 1 + min(curr[j-1], prev[j-1], prev[j])
			}
		}
		for _, v := range curr {
			r += v
		}
		copy(prev, curr)
		for j := range curr {
			curr[j] = 0
		}
	}
	return r
}

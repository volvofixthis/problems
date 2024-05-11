package first

import (
	"math"
)

func reverse(x int) int {
	var result int64
	xu := x
	if xu < 0 {
		xu = xu * -1
	}
	for xu > 0 {
		digit := xu % 10
		xu = xu / 10
		result = result*10 + int64(digit)
	}
	if x < 0 {
		result = result * -1
	}
	if result < int64(math.Pow(-2, 31)) || result > int64(math.Pow(2, 31))-1 {
		result = 0
	}
	return int(result)
}

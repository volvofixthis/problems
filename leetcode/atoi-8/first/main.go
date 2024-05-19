package first

import (
	"math"
)

var low = int(math.Pow(-2.0, 31.0))
var high = int(math.Pow(2.0, 31.0)) - 1

func myAtoi(s string) int {
	sign := 1
	stg := 0
	n := 0
	for i := 0; i < len(s); i++ {
		if stg == 0 && s[i] == '-' {
			sign = -1
			stg = 1
			continue
		} else if stg == 0 && s[i] == '+' {
			stg = 1
			continue
		} else if stg == 0 && s[i] == ' ' {
			continue
		} else if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
			if n*sign < low {
				return low
			}
			if n*sign > high {
				return high
			}
			stg = 1
			continue
		} else {
			break
		}
	}
	return n * sign
}

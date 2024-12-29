package main

func maximum69Number(num int) int {
	st := 10000
	for st >= 10 {
		md := num % st
		st = st / 10
		digit := md / st
		if digit == 6 {
			num += st + st + st
			break
		}
	}
	return num
}

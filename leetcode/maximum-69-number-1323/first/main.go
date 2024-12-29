package main

import "strconv"

func maximum69Number(num int) int {
	s := []byte(strconv.Itoa(num))
	for i := 0; i < len(s); i++ {
		if s[i] == '6' {
			s[i] = '9'
			break
		}
	}
	r, err := strconv.Atoi(string(s))
	if err != nil {
		return 0
	}
	return r
}

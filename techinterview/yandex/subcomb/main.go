package main

import "fmt"

/*
Найти вхождение подстроки в строке. Строки считаются равными, если из одной подстроки можно получить
другую перестановкой символов. В противном случае нужно вернуть -1.
"reebok", "bee" -> 1
*/

func findSubPos(s, sub string) int {
	if len(s) < len(sub) {
		return -1
	}
	cc := map[byte]int{}
	l := len(sub)
	update := func(b byte, v int) {
		nv := cc[b] + v
		if nv == 0 {
			delete(cc, b)
			return
		}
		cc[b] = nv
	}
	for i := range l {
		update(sub[i], 1)
		update(s[i], -1)
	}
	matched := func() bool {
		if len(cc) != 0 {
			return false
		}
		return true
	}
	if matched() {
		return 0
	}
	for i := 1; i < len(s)-l; i++ {
		update(s[i-1], 1)
		update(s[i+l-1], -1)
		if matched() {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findSubPos("reebok", "bee"))
}

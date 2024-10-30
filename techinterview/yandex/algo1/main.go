package main

import (
	"fmt"
	"strings"
)

func normalize(s string) string {
	stack := make([]string, 0, 10)
	part := []rune{}
	process := func() {
		switch string(part) {
		case "../", "..":
			if len(stack) == 1 && stack[0] == "/" {
				break
			} else if len(stack) > 0 && stack[len(stack)-1] != "/" {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, string(part))
			}
		case "./", ".":
		case "/":
			if len(stack) > 0 {
				break
			}
			fallthrough
		default:
			stack = append(stack, string(part))
		}
	}
	for i, r := range s {
		fmt.Println(i)
		part = append(part, r)
		if r == '/' {
			process()
			part = part[:0]
			continue
		}
	}
	if len(part) > 0 {
		process()
	}
	if len(stack) > 0 {
		last := stack[len(stack)-1]
		if last != "/" && last[len(last)-1] == '/' {
			stack[len(stack)-1] = last[:len(last)-1]
		}
	}
	return strings.Join(stack, "")
}

func main() {
	fmt.Println(normalize("/etc/../kus"))
}

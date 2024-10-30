package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	as := len(a)
	for i := 0; i <= as; i++ {
		fmt.Println(i)
		a = a[:i]
		fmt.Println(len(a), cap(a))
		fmt.Println(a)
	}
}

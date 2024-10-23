package main

import "fmt"

var m map[any]any = make(map[any]any)

func Add[T any](v T) {
	type K struct {
		a *string
	}
	m[K{}] = v
}

func main() {
	Add(123)
	Add(true)
	fmt.Println(len(m))
	var a []int
	if a == nil {
		fmt.Println("kus")
	}
	var c interface{} = (*int)(nil)
	if c == nil {
		fmt.Printf("kus 2")
	}
}

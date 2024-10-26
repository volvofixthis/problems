package main

import "fmt"

func main() {
	var numbers []*int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, &i)
	}
	*numbers[0] = *numbers[0] * 2
	for _, v := range numbers {
		fmt.Println(*v)
	}
}

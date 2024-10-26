package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := make([]int, 3, 5)
	fix(numbers)
	numbers = numbers[:4]
	sr(numbers)
	fmt.Println(numbers)
}

func fix(numbers []int) {
	numbers[0] = 3
	numbers[1] = 2
	numbers[2] = 1
	numbers = append(numbers, 4)
}

func sr(numbers []int) {
	sort.Ints(numbers)
}

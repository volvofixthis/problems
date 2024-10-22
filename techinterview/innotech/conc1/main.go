package main

import (
	"fmt"
)

func main() {
	for v := range square(generator()) {
		fmt.Println(v)
	}
}

func generator() <-chan int {
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	return ch1
}

func square(ch <-chan int) <-chan int {
	ch2 := make(chan int)
	go func() {
		for v := range ch {
			ch2 <- v * v
		}
		close(ch2)
	}()
	return ch2
}

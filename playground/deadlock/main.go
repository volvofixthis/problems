package main

import (
	"fmt"
	"math"
)

func main() {
	ch := make(chan int)

	go func() {
		// time.Sleep(2 * time.Second) // Simulate a delay
		v := int64(2)
		for i := int32(0); i < math.MaxInt32; i++ {
			v *= 2
		}
		ch <- 1
	}()

	<-ch // Will wait until the goroutine sends a value
	fmt.Println("kus")
}

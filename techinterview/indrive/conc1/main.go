package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("one")
	time.Sleep(1 * time.Second)
	go func() {
		fmt.Println("two")
	}()
	go func() {
		fmt.Println("three")
	}()
}

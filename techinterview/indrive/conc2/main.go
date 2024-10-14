package main

import (
	"fmt"
	"sync"
)

func main() {
	fetchConcurrently(getData(), 4)
}

func fetchConcurrently(urls []string, n int) {
	c := make(chan string)
	wg := sync.WaitGroup{}
	for range n {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				if v, ok := <-c; ok {
					if !checkUrl(v) {
						fmt.Printf("Kus %s!\n", v)
					}
					continue
				}
				break
			}
		}(&wg)
	}
	for _, u := range urls {
		c <- u
	}
	close(c)
	wg.Wait()
}

func getData() []string {
	return []string{"test1", "test2", "test3", "test4", "test5", "test6"}
}

func checkUrl(_ string) bool {
	return false
}

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	processConc(generate(), 4)
}

func process(s string) {
	fmt.Println(s)
}

func processConc(ch <-chan string, wn int) {
	wg := sync.WaitGroup{}
	for i := 0; i < wn; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for v := range ch {
				process(v)
			}
		}(&wg)
	}
	wg.Wait()
}

func generate() chan string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	ch := make(chan string, 4)
	go func(ctx context.Context) {
		defer cancel()
		i := 0
		for {
			i++
			select {
			case <-ctx.Done():
				close(ch)
				return
			case ch <- fmt.Sprintf("%d", i):
			}
		}
	}(ctx)
	return ch
}

package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func Process() int {
	time.Sleep(1 * time.Second)
	return 12
}

func ProcessWrapper(ctx context.Context) (int, error) {
	ch := make(chan int)
	go func() {
		ch <- Process()
		close(ch)
	}()
	select {
	case v := <-ch:
		return v, nil
	case <-ctx.Done():
		return 0, errors.New("Timeouted")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	fmt.Println(ProcessWrapper(ctx))
}

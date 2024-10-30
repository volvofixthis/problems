package main

/*
Есть массив задач с методом Run() error
Реализовать функцию func execute(tasks []Task) []error, которая запускает каждую задачу в своей горутине и
возвращает массив (слайс!!!) ошибок

*/

import (
	"fmt"
	"sync"
)

const workerNumber = 4

type Task interface {
	Run() error
}

type EmailTask struct {
}

func (et *EmailTask) Run() error {
	return nil
}

func main() {
	errs := execute(
		[]Task{
			new(EmailTask),
			new(EmailTask),
			new(EmailTask),
			new(EmailTask),
		},
	)
	fmt.Println(errs)
}

func execute(tasks []Task) []error {
	ch1 := make(chan Task, len(tasks))
	ch2 := make(chan error, len(tasks))
	wg := sync.WaitGroup{}
	for _, t := range tasks {
		ch1 <- t
	}
	close(ch1)
	for range workerNumber {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range ch1 {
				ch2 <- t.Run()
			}
		}()
	}
	wg.Wait()
	close(ch2)
	errors := make([]error, 0, len(tasks))
	for {
		select {
		case err, ok := <-ch2:
			if !ok {
				goto out
			}
			errors = append(errors, err)
		}
	}
out:
	return errors
}

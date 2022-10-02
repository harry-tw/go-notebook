package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func intGenerator(done <-chan interface{}) <-chan int {
	rand.Seed(time.Now().Unix())
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- rand.Intn(100):
			}
		}
	}()
	return out
}
func multiply(done <-chan interface{}, in <-chan int, by int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for val := range in {
			select {
			case <-done:
				return
			case out <- val * by:
			}
		}
	}()
	return out
}
func fanIn(done <-chan interface{}, ins ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	aggregate := func(in <-chan int) {
		defer wg.Done()
		for val := range in {
			select {
			case <-done:
				return
			case out <- val:
			}
		}
	}
	wg.Add(len(ins))
	for _, in := range ins {
		go aggregate(in)
	}
	go func() {
		wg.Wait()
		defer close(out)
	}()
	return out
}

func getResult(done chan interface{}, in <-chan int, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case out <- <-in:
			}
		}
	}()
	return out
}

func fanOut(num int, worker <-chan int) []<-chan int {
	workers := make([]<-chan int, num)
	for i := 0; i < num; i++ {
		workers[i] = worker
	}
	return workers
}

func main() {
	done := make(chan interface{})
	defer close(done)

	generator := intGenerator(done)
	workers := fanOut(runtime.NumCPU(), multiply(done, generator, 2))
	for result := range getResult(done, fanIn(done, workers...), 10) {
		fmt.Println(result)
	}
	return
}

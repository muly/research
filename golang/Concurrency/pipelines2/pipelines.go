// Go Concurrency Patterns: Pipelines and cancellation

package pipelines

import (
	"sync"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {

	wg := sync.WaitGroup{}

	out := make(chan int)

	//
	output := func(c <-chan int) {
		for i := range c {
			out <- i
		}
		wg.Done()

	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	time.Sleep(600000)

	go func() {
		//wg.Wait()
		close(out)
	}()
	return out
}

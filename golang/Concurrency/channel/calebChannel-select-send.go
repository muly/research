package main

import (
	"fmt"
	"time"
)

func main() {

	c1, c2 := make(chan int), make(chan int)
	timeout := time.After(time.Nanosecond * 900)
	go func() {
		for {
			select {
			case c1 <- 1:
				//time.Sleep(500 * time.Millisecond)
			case c2 <- 2:
				//time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for {
		select {
		case i := <-c1:
			fmt.Println(i)
		case i := <-c2:
			fmt.Println(i)
		case <-timeout:
			fmt.Println("TimedOut")
			return
		}
	}

}

// Fan-in

// Generator pattern

package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan time.Duration)
	timeout := time.After(time.Millisecond * 4000)
	go f1(c)

	for {
		select {
		case msg := <-c:
			fmt.Println(msg)
		case <-timeout: // time.After(time.Millisecond * 4000):
			fmt.Println("TimedOut")
			return
		}
	}

}

func f1(c chan time.Duration) {
	for {
		for i := 0; i < 5; i++ {
			d := time.Millisecond * time.Duration(i*100)
			time.Sleep(d)
			c <- d
		}
	}
}

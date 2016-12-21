// Generator pattern

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := printer("Srini ")

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("you are boring, I'm leaving")

}

// channel generator
func printer(msg string) <-chan string {
	//func printer(msg string) (c <-chan string) { // ERROR
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- "Hello " + msg + "-" + strconv.Itoa(i)
			time.Sleep(time.Second)

		}
	}()
	return c
}

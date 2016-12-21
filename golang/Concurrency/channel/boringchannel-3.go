// Generator pattern

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	srini := printer("Srini")
	neil := printer("Neil")

	for i := 0; i < 5; i++ {
		fmt.Println(<-srini)
		fmt.Println(<-neil)
	}

	fmt.Println("you are boring, I'm leaving")

}

// channel generator
func printer(msg string) <-chan string {
	//func printer(msg string) (c <-chan string) { // ERROR
	c := make(chan string)
	go func() {
		for i := 0; i < 7; i++ {

			c <- "Hello " + msg + "-" + strconv.Itoa(i)

			switch msg {
			case "Srini":
				time.Sleep(time.Second)
			case "Neil":
				time.Sleep(time.Second * 4)
			default:
				time.Sleep(time.Second)
			}
		}
	}()
	return c
}

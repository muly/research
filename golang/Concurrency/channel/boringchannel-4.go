// Fan-in

// Generator pattern

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := fanIn(printer("Srini"), printer("Neil"))

	for i := 0; i < 14; i++ {
		fmt.Println(<-c)
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
				time.Sleep(time.Second * 4)
			case "Neil":
				time.Sleep(time.Second)
			default:
				time.Sleep(time.Second)
			}
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {

	c := make(chan string)

	go func() {
		for { //Q: why is this infinite for loop required?
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c

}

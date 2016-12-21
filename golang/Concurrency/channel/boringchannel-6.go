package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	srini := printer("Srini")
	neil := printer("Neil")

	for i := 0; i < 14; i++ {
		select {
		case v1 := <-srini:
			fmt.Println(v1)
		case v2 := <-neil:
			fmt.Println(v2)
			//default: fmt.Println("no message")
		}
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

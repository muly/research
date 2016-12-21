package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	c := make(chan string)
	go printer("Srini ", c)

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("you are boring, I'm leaving")

}

func printer(msg string, c chan string) {
	for i := 0; i < 10; i++ {
		c <- "Hello " + msg + "-" + strconv.Itoa(i)
		time.Sleep(time.Second)

	}

}

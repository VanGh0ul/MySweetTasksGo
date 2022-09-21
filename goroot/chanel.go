package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
func printer(c <-chan string) {
	for {
		fmt.Println(<-c)
		amt := time.Duration(rand.Intn(8))
		time.Sleep(time.Millisecond * amt)
	}
}
func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

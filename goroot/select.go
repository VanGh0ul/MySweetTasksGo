package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	c3 := make(chan string, 1)
	c4 := make(chan string, 1)
	c5 := make(chan string, 1)
	c6 := make(chan string)
	c7 := make(chan string)
	c8 := make(chan string)
	c9 := make(chan string)
	c10 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			amt := time.Duration(rand.Intn(1000))
			time.Sleep(time.Millisecond * amt)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"

		}
	}()
	go func() {
		for {
			c3 <- "from 3"
		}
	}()
	go func() {
		for {
			c4 <- "from 4"
		}
	}()
	go func() {
		for {
			c5 <- "from 5"
		}
	}()
	go func() {
		for {
			c6 <- "from 6"
		}
	}()
	go func() {
		for {
			c7 <- "from 7"
		}
	}()
	go func() {
		for {
			c8 <- "from 8"
		}
	}()
	go func() {
		for {
			c9 <- "from 9"
		}
	}()
	go func() {
		for {
			c10 <- "from 10"
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case msg3 := <-c3:
				fmt.Println("Message 3", msg3)
			case msg4 := <-c4:
				fmt.Println("Message 4", msg4)
			case msg5 := <-c5:
				fmt.Println("Message 5", msg5)
			case msg6 := <-c6:
				fmt.Println("Message 6", msg6)
			case msg7 := <-c7:
				fmt.Println("Message 7", msg7)
			case msg8 := <-c8:
				fmt.Println("Message 8", msg8)
			case msg9 := <-c9:
				fmt.Println("Message 9", msg9)
			case msg10 := <-c10:
				fmt.Println("Message 10", msg10)

			case <-time.After(time.Second):
				fmt.Println("timeout")

			default:
				fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

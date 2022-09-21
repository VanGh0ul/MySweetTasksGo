package main

import "fmt"

func funcIter(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		go funcIter(i)
	}
	var input string
	fmt.Scanln(&input)
}

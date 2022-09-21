package main

import "fmt"

func main() {
	fmt.Println("Строки:")
	fmt.Println("5 / 2 = ", 5.0/2.0)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[6])
	fmt.Println("Hello" + "World")
	fmt.Println("Литералы:")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

package main

import "fmt"

func swap(x, y *int) {

*x, *y = *x, *x
}
func main() {
	x := 1
	y := 2
	fmt.Println("before x =", x, "before y =", y)
	swap(&x, &y)
	fmt.Println("after x =", x, "after y =", y)
}
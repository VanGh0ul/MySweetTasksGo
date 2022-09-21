package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))
	c.x = 10
	c.y = 5
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

}

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}


type distance struct {
	f1, f2, f3, f4 float64
}

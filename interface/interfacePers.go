package main

import (
	"fmt"
)

func main() {
	
	r := 1
	c := 2
	fmt.Println(totalArea(&c, &r))
}

type Shape interface {
    area() float64
}

type MultiShape struct {
    shapes []Shape
}
func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

package main

import (
	"fmt"
)

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nexEven := makeEvenGenerator()
	fmt.Println(nexEven())
	fmt.Println(nexEven())
	fmt.Println(nexEven())
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	x1 := 9

	i := sort.Search(len(x), func(i int) bool { return x1 <= x[i] })
	if i < len(x) && x[i] == x1 {
		fmt.Printf("Найдено %s по индексу %d в %v.\n", x1, i, x)
	} else {
		fmt.Printf("Не найдено %s в %v.\n", x1, x)
	}

}

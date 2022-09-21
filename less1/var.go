package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")

	var input float64

	var C float64

	var F float64

	fmt.Scanf("%f", &input)
	//фарингейт -> цельсий
	fmt.Println("фарингейт -> цельсий")
	C = (input - 32) * 5 / 9
	fmt.Println(C)
	//футы -> метры
	fmt.Println("футы -> метры")
	F = input * 0.3048
	fmt.Println(F)
}

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		//task 1 Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. (3, 6, 9, …).
		if i%3 == 0 {
			fmt.Println(i)
		}

	}
	for i := 1; i <= 100; i++ {
		//task 2 Напишите программу, которая выводит числа от 1 до 100. Но для кратных трём нужно вывести «Fizz» вместо числа, для кратных пяти - «Buzz», а для кратных как трём, так и пяти — «FizzBuzz».
		if i%3 == 0 {
			fmt.Println(i, "Fizz")
		}

		if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}

		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		}
	}
}

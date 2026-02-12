package main

import "fmt"

func main () {
	/*
		If construct
	*/

	var num1, num2 int = 45, 6

	if num2 == 0 { // You can omit the parentheses
		fmt.Printf("Non puoi dividere per 0\n")

		return
	}

	quotient := num1 / num2
	fmt.Printf("La divisione intera tra %d e %d fa %d\n", num1, num2, quotient)

	/*
		For construct
	*/

	// 1st method to use a for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}

	// 2nd method to use a for loop
	for i := range 10 {
		fmt.Printf("%d\n", i)
	}
}
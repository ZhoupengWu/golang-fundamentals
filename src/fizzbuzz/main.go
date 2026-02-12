package main

import (
	"fmt"
	"strconv"
)

func main () {
	fmt.Printf("===== FIZZBUZZ =====\n")

	for i := range 20 {
		fmt.Printf("Test %d: %s\n", i, fizzbuzz(i))
	}
}

func fizzbuzz (input int) string {
	switch {
	case (input % 3 == 0) && (input % 5 == 0):
		return "FizzBuzz"
	case input % 3 == 0:
		return "Fizz"
	case input % 5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(input)
	}
}
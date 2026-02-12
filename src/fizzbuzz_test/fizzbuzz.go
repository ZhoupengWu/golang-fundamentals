package fizzbuzztest

import "strconv"

func Fizzbuzz (input int) string {
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
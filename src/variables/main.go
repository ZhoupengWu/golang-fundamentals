package main

import "fmt"

func main () {
	/*
		In the 2nd and 3rd method you must assign a value to the variable
	*/
	var stato string = "Italia" // 1st method to declare and assign a value specifying the type of variable
	var car string = "Lamborghini"
	cheap_car := "BMW" // 2nd method to declare and assign a value without specifying the type of variable
	var school = "Galileo Galilei" // 3rd method to declare and assign a value without specifying the type of variable
	_ = school // This tells Go that the variable is used

	fmt.Printf("La mia automobile è una %s e vivo in %s\n", car, stato)
	fmt.Printf("La mia amica ha una %s\n", cheap_car)

	var friend_age, major_age int = 17, 18
	is_adult := friend_age > major_age

	fmt.Printf("La mia amica è maggiorenne? %t\n", is_adult)

	/*
		In go, we have various types

		- Integer: int, int8, int16, int32, int64
		- Integer unsigned: uint, uint8, uint16, uint32, uint64
		- Float: float32, float64
		- Boolean: bool
		- String: string
		- Special types: - complex64, complex128 (for complex numbers)
						 - byte (alias for uint8)
						 - rune (alias for int32, used to represent a Unicode character)
						 - uintptr (used to store a generic pointer)
	*/

	/*
		To convert from a type to another type, we use Type(name_variable)

		- Ex1: a := int(4.54)
		- Ex2: a := 45; b := float64(a)
	*/

	/*
		Const: a variable that can't modify. It's defined at compile-time
	*/
	const PI = 3.1429

	/*
		Iota: a counter that start from 0 and increment by 1 for each const of the block. Only used in a const block
		You can create expressions with iota
	*/
	const (
		MONDAY = iota + 1
		TUESDAY
		WEDNESDAY
		THURSDAY
		FRIDAY
		SATURDAY
		SUNDAY
	)

	const (
		Bi = 8 << (iota * 10)
		KBi
		MBi
		GBi
		TBi
	)

	fmt.Printf("1 terabit sono %d\n", TBi)
}
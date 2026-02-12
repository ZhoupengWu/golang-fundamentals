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
}
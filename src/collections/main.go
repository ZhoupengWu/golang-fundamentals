package main

import (
	"fmt"
	"sort"
)

func main () {
	/*
		Collections: array
	*/

	// Arrays are declared and assigned in the same way as slices
	// 1st method to declare an array and assign values
	var marks [5]int
	marks[0] = 5
	marks[1] = 8
	marks[2] = 10
	marks[3] = 4
	marks[4] = 7

	// 2nd method to declare an array and assign values
	var balls [2]int = [2]int{104, 401}

	// 3rd method to declare an array and assign values
	my_iphones := [3]int{17, 16, 15}

	// 4th method to declare an array and assign values
	name_girl := [...]string{"Ilaria", "Gaia", "Viola", "Sofia"}

	// 5th method to declare an array and assign values
	prezzi := [...]int{1: 45, 2: 99}

	fmt.Printf("Array: %v\n", marks)
	fmt.Printf("Array: %v\n", balls)
	fmt.Printf("Array: %v\n", my_iphones)
	fmt.Printf("Array: %v\n", name_girl)
	fmt.Printf("Array: %v\n", prezzi)

	// We can access a subset of an array ([start:end] --- start inclusive and end exclusive)
	girl_friends := name_girl[1:3]
	girl_friends = append(girl_friends, "aa", "bb")

	// We can check with = if two arrays have the same length and the same elements in the same order
	fmt.Println(name_girl == [4]string(girl_friends))

	// Array has a various functions
	// sort, using the standard library of go
	// copy: it takes two arguments (dest array, source array) and returns the number of elements copied
	// Others functions useful: sort.Search, sort.Reverse

	sort.Ints(my_iphones[:])
	copied := copy(prezzi[:], balls[:])
	_ = copied

	// You can declare array of a type for example [length]type

	/*
		Collections: slice (array dinamici)
	*/

	friends := []string{"Pablo", "Marco", "Keplero"}
	fmt.Printf("I miei amici ora: %v\n", friends)

	friends = append(friends, "Chiara")
	fmt.Printf("I miei amici dopo: %v\n", friends)

	// Other methods to declare an array and assign values
	var a []int // Value nil
	var b []int = []int{} // Value []
	c := make([]int, 5)
	d := make([]int, 0, 10) // With 3rd argument you allocate an array with size 10 without initialising the values to 0
	e := make([]int, 5, 10) // With 3rd argument you allocate an array with size 10 initialising the first 5 values to 0
	var f []int = make([]int, 5)
	g := marks[2:4] // If capacity is not enough, the variable points a different memory's address otherwise the address of the copied array

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e, f, g)

	/*
		Collections: map (dict as in python)
	*/

	// Maps are declared and assigned in the same way as slices
	students := make(map[string]string, 3)
	students["Angela"] = "4IC"
	students["Paola"] = "4IA"
	students["Andrea"] = "4IB"
	students["Luca"] = "4MA"
	students["Maria"] = "4TA"

	var students2 map[string]string = map[string]string{"Pablo": "5IG", "Anna": "9MA"}
	students3 := map[string]string{"Lara": "7TB", "Alice": "1IA"}

	fmt.Printf("%v\n%v\n", students2, students3)

	for name, class := range students {
		fmt.Printf("%s sta nella classe %s\n", name, class)
	}

	/*
		String can be considered as a slice of byte or rune

		There are 2 types of string: - String literals (using "")
									 - Raw string literals (using ``) --> for multiline strins
	*/
}
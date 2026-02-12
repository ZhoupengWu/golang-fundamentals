package main

import (
	"fmt"
	"strings"
)

func main () {
	var num int = 19

	fmt.Printf("Valore numero: %d ----- Indirizzo: %p\n", num, &num)

	var ptr *int = &num

	fmt.Printf("Valore puntatore: %d ----- Indirizzo: %p\n", *ptr, ptr)

	num = 99

	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Valore numero: %d ----- Indirizzo: %p\n", num, &num)
	fmt.Printf("Valore puntatore: %d ----- Indirizzo: %p\n", *ptr, ptr)

	*ptr = 104

	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Valore numero: %d ----- Indirizzo: %p\n", num, &num)
	fmt.Printf("Valore puntatore: %d ----- Indirizzo: %p\n", *ptr, ptr)

	/*
		Allocate an array of size 3 with nil value (you have to use new(type) to allocate *type value)
	*/
	ptr_a := make([]*int, 3)
	ptr_a[0] = new(int)
	ptr_a[1] = new(int)
	ptr_a[2] = new(int)
	*ptr_a[0] = 1
	*ptr_a[1] = 2
	*ptr_a[2] = 3

	for _, v := range ptr_a {
		fmt.Println(strings.Repeat("-", 50))
		fmt.Printf("Valore puntatore: %d ----- Indirizzo: %p\n", *v, v)
	}
}
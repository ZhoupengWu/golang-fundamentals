package main

import (
	"fmt"
	"strings"
)

type player struct {
	name string
	surname string
	age int
	country string
}

func main () {
	player1 := player{
		name: "Paolo",
		surname: "Hots",
		age: 33,
		country: "USA",
	}

	// 2nd method to declare a variable
	var player2 player = player{
		name: "Ciccio",
		surname: "Gamer",
		age: 26,
		country: "Italy",
	}

	fmt.Printf("Player 1: %v\n", player1)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Player 2: %v\n", player2)
}
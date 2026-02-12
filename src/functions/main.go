package main

import (
	"errors"
	"fmt"
	"functions/repo"
)

type player struct {
	name string
	surname string
	age int
	country string
}

// You can declare methods for all custom type, not only for struct, for example type vehicle string
type money float64

func main () {
	player1 := player{
		name: "Paolo",
		surname: "Hots",
		age: 33,
		country: "USA",
	}

	printPlayer(player1)
	player1.printPlayerMethod()

	num1, num2 := 7, 8
	result, err := divide(num1, num2)

	if err != nil {
		fmt.Printf("ERRORE: %s\n", err)
	} else {
		fmt.Printf("Risultato: %d\n", result)
	}


	ford := repo.Car{
		Name: "Mario",
		Cost: 56000,
	}

	ford.DisplayInfo()

	/*
	  *************************
	  *************************
	  *************************
	  *************************
	  *************************
	*/

	// Methods of a custom type
	my_wallet := money(5000)

	my_wallet.print()
	my_wallet.addMoney()
	my_wallet.print()

	// Variadic function
	fmt.Printf("PIL of Italy: %d$\n", pil(500000, 400000, 230000, 890000, 1400000))
	fmt.Printf("PIL of France: %d$\n", pil(590000, 360000, 950000, 1190000, 1650000))
	fmt.Printf("PIL of USA: %d$\n", pil(1800000, 2400000, 2900000, 1890000, 14000000))

	// Named return values
	_, _ = lalalala()

	// Functions as variables
	getAura := func (aura int) int {
		return aura + 1000
	}

	fmt.Printf("Your aura now is %d pt\n", getAura(9000))

	// Closure function
	nextPos := count()

	for range 5 {
		fmt.Printf("Pos: %d\n", nextPos())
	}

	first_machine := bonusWin(8)

	for i := range 10 {
		fmt.Printf("Your bonus is %d\n", first_machine(i))
	}
}

func printPlayer (player player) {
	fmt.Printf("Name: %s --- Surname: %s --- Age: %d --- Country: %s\n", player.name, player.surname, player.age, player.country)
}

func (p *player) printPlayerMethod () {
	fmt.Printf("Name: %s --- Surname: %s --- Age: %d --- Country: %s\n", p.name, p.surname, p.age, p.country)
}

func divide (a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	return a / b, nil
}



// Methods of money

func (m *money) addMoney () {
	*m += 5000
}

func (m money) print () {
	fmt.Printf("In your wallet you have %.2f$\n", m)
}

// Variadic function

func pil (moneys ...int) int {
	total := 0

	for _, money := range moneys {
		total += money
	}

	return total
}

// Named return values
func lalalala () (a, b string) {
	a = "Apple"
	b = "Melon"

	return
}

// Closure function
func count () func() int {
	c := 0

	return func() int {
		c++

		return c
	}
}

func bonusWin (initialBonus int) func(nextBonus int) int {
	return func(nextBonus int) int {
		return initialBonus * nextBonus
	}
}
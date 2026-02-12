package repo

import "fmt"

type Car struct {
	Name string
	Cost int
}

func (c Car) DisplayInfo () {
	fmt.Printf("La tua %s costa %d$\n", c.Name, c.Cost)
}
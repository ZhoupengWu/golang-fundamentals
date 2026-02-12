package main

import "fmt"

func main () {
	workers := []worker{
		&softwareDeveloper{
			name: "Giovanni",
			surname: "Manetto",
			programmingLanguages: []string{
				"GO",
				"Java",
				"JavaScript",
			},
		},
		&softwareDeveloper{
			name: "Marco",
			surname: "Ladrone",
			programmingLanguages: []string{
				"Python",
				"Php",
				"C",
			},
		},
	}

	for _, w := range workers {
		fmt.Println(w.work())
		w.changeName("COGLIONE")
		fmt.Println(w.work())
	}
}

// Esistono le empty interfaces
type worker interface {
	work() string
	changeName(newName string)
}

type softwareDeveloper struct {
	name string
	surname string
	programmingLanguages []string
}

func (s *softwareDeveloper) work () string {
	return fmt.Sprintf("%s %s is working", s.name, s.surname)
}

func (s *softwareDeveloper) changeName (newName string) {
	s.name = newName
}
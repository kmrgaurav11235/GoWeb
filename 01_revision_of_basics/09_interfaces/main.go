package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
} /*
The above code basically says:
	1. Our program now has a new type called 'Animal'.
	2. If you are a type in this program with a function named 'Says()' that returns string and
		another named 'NumberOfLegs()' which also returns string, then you are now an honorary
		member of the type 'Animal'.
Note that interfaces are implicit, i.e. we don't have to manually say that our custom types satisy
the interface.
*/

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Mickey",
		Breed: "Shiba",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Harambe",
		Color:         "Black",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This Animal says", a.Says(), "and has", a.NumberOfLegs(), "legs.")
}

// In Go, most receivers are of pointer types. This makes things faster and is recommended in Go docs.
func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Ugh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}

/*
More points about Interfaces:
* Interfaces are not generic types. Go (famously) does not have generic types.
* Interfaces are a contract to help us manage types.
*/

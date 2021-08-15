package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// Create Map: The make built-in function allocates and initializes an object of type slice, map, or chan.
	myMap := make(map[string]string)

	myMap["Narrator"] = "Robinson Crusoe"
	myMap["Servant"] = "Friday"

	log.Println(myMap)

	anotherMap := make(map[string]User)

	robinson := User{
		FirstName: "Robinson",
		LastName:  "Crusoe",
	}

	father := User{
		FirstName: "Kreutznaer",
	}

	anotherMap["First"] = robinson
	anotherMap["Second"] = father

	log.Println(anotherMap["First"], anotherMap["Second"])

	// Slices are lists
	var mySlice []string

	mySlice = append(mySlice, "The widow")
	mySlice = append(mySlice, "The Spaniard")
	mySlice = append(mySlice, "Portuguese Sea Captain")

	log.Println(mySlice)

	// Another way to create slice
	intSlice := []int{3, 1, 11, -8, 4, 32}
	log.Println("Slice", intSlice)

	sort.Ints(intSlice)
	log.Println("Sorted Slice", intSlice)

	// Only the first 3 numbers from "intSlice"
	log.Println("First 3 numbers", intSlice[0:3])
}

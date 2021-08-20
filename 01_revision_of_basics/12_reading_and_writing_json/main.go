package main

import (
	"encoding/json"
	"fmt"
)

// A type created on the basis of myJson
// The json fields tell go how to serialize/deserailize this
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]` // Multi line string can be declared using ``

	// 1. Read json into a struct

	var unmarshalled []Person // Since myJson is a slice of Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by 2nd arg.

	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	fmt.Printf("Unmarshalled: %v\n", unmarshalled)

	// 2. Write json from a struct

	m1 := Person{
		FirstName: "Wally",
		LastName:  "West",
		HairColor: "red",
		HasDog:    false,
	}

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "Black"
	m2.HasDog = true

	var mySlice []Person
	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling json", err)
	}

	fmt.Println("Marshalled:", string(newJson))
}

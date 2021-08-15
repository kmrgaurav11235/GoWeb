package main

import "log"

type User struct {
	FirstName string
	LastName string
}

func main() {
	// Create Map: The make built-in function allocates and initializes an object of type slice, map, or chan.
	myMap := make(map[string]string)

	myMap["Narrator"] = "Robinson Crusoe"
	myMap["Servant"] = "Friday"

	log.Println(myMap)

	anotherMap := make(map[string]User)

	robinson := User {
		FirstName: "Robinson",
		LastName: "Crusoe",
	}

	father := User {
		FirstName: "Kreutznaer",
	}

	anotherMap["First"] = robinson
	anotherMap["Second"] = father

	log.Println(anotherMap["First"], anotherMap["Second"])
}

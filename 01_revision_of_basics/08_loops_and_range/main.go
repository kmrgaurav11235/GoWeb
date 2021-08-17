package main

import "fmt"

func main() {
	// Simple for-loop
	for i := 1; i <= 10; i++ {
		fmt.Println("5 *", i, "=", (5 * i))
	}

	// for-loop with range
	travellers := []string{"Phileas Fogg", "Jean Passepartout", "Detective Fix", "Aouda"}

	// for i, traveller := range travellers {
	// 	fmt.Println("Traveller", i+1, ":", traveller)
	// }

	// Use _ if you don't intend to use the value
	for _, traveller := range travellers {
		fmt.Println(traveller)
	}

	// Range over maps
	characters := make(map[string]string)

	characters["Professor"] = "Otto Lidenbrock"
	characters["Nephew"] = "Axel"
	characters["Guide"] = "Hans Bjelke"
	characters["Goddaughter"] = "Gretchen"

	// map is accessed in no particular order
	for what, who := range characters {
		fmt.Println(what, ":", who)
	}

	// You can also range over other types. e.g. a string
	var firstLine = "Call me Ishmael"

	for i, l := range firstLine {
		fmt.Println(i, ":", l) /* In Go, a string just a slice of bytes. So, this will print index and the
		numbers corresponding to each of the letter in the string.*/
	}
	// Also, in Go, strings are immutable

	// We can also range over slices of custom objects
	type User struct {
		FirstName string
		LastName  string
		Email     string
	}

	var users []User
	users = append(users, User{FirstName: "Captain", LastName: "Ahab", Email: "captain.ahab@pequod.org"})
	users = append(users, User{FirstName: "Ishmael", LastName: "", Email: "ishmael@pequod.org"})
	users = append(users, User{FirstName: "Moby", LastName: "Dick", Email: "moby.dick@badass.org"})
	users = append(users, User{FirstName: "Father", LastName: "Mapple", Email: "father.mapple@church.org"})

	for _, user := range users {
		fmt.Println(user.FirstName, user.LastName, user.Email)
	}
}

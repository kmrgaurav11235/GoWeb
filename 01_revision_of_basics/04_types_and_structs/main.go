package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

// In Go, a variable or Method starting with a small letter cannot be accessed outside the package.
// But, a variable or Method starting with a capital letter can be accessed outside the package.

func main() {
	user := User{
		FirstName:   "Captain",
		LastName:    "Nemo",
		PhoneNumber: "1-767–1799", // In Go, the last line has a comma too
	}

	log.Println(user.FirstName, user.LastName, user.BirthDate)
}

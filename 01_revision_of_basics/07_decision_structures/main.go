package main

import "log"

func main() {
	// if - else if - else
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is true")
	} else {
		log.Printf("isTrue is not true")
	}

	myNum := 100
	isTrue = false

	if myNum > 50 && !isTrue {
		log.Println("myNum is greater than 50 and isTrue is set to false.")
	} else if myNum > 50 && isTrue {
		log.Println("myNum is greater than 50 and isTrue is set to true.")
	} else {
		log.Println("myNum is less than or equal to 50 and isTrue is set to ", isTrue)
	}

	// switch - No need for break in Go.
	name := "Arne Saknussemm"

	switch name {
	case "Otto Lidenbrock":
		log.Println("A hot-tempered geologist with radical ideas")
	case "Axel":
		log.Println("Lidenbrock's nephew, a young student whose ideas are more cautious.")
	case "Arne Saknussemm":
		log.Println("Discovered a passage to the center of the earth via the crater of Snæfellsjökull in Iceland")
	case "Hans Bjelke":
		log.Println("Icelandic eiderduck hunter who hires on as their guide; resourceful and imperturbable.")
	default:
		log.Println("Character not available.")
	}

}

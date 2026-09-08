package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age // Get the address

	fmt.Println("Age: ", age)
	editAgeToAdultYears(agePointer)
	fmt.Println("Age: ", *agePointer) // Get the value
}

func editAgeToAdultYears(age *int) {
	*age -= 18
}

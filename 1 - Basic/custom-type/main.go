package main

import "fmt"

// Assign an alias
type str string

func (text str) log() {
	// Method
	fmt.Println(text)
}

func main() {
	var name str

	name = "Gabriel"

	name.log()

}

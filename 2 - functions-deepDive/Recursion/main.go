package main

import (
	"fmt"
)

func main() {

	value := 5

	fmt.Println(factorial(value))
}

func factorial(value int) int {

	if value == 0 {
		return 1
	}

	return value * factorial(value-1)

}

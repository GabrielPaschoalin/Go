package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	fmt.Println(sumup(numbers))
	fmt.Println(variadicFunc(3, 1, 2, 3, -6, 200))

	anotherSum := variadicFunc(2, numbers...) // Transforma o slice em uma lista de parâmetros
	fmt.Println(anotherSum)
}

func sumup(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func variadicFunc(mult int, numbers ...int) int {
	// Work with any amount of parameters
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum * mult
}

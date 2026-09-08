package main

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 10}
	fmt.Println(tranformNumbers(&numbers, selectType(&numbers)))
	fmt.Println(tranformNumbers(&moreNumbers, selectType(&moreNumbers)))

}

func tranformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func selectType(numbers *[]int) func(int) int {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

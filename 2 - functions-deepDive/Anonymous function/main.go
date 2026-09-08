package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// Anonymous function
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformed)

	// Factory function / closure
	double := createTransformer(2)
	triple := createTransformer(3)
	fmt.Println(transformNumbers(&numbers, double))
	fmt.Println(transformNumbers(&numbers, triple))

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	// Factory function
	return func(number int) int {
		return number * factor
	}
}

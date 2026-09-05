package main

import "fmt"

func main() {

	printSomething("-------------------")
	printSomething(1)
	printSomething(2.1)

	typeVal("-------------------")
	typeVal(1)
	typeVal(2.1)

	result := add(1, 2.0)
	fmt.Println(result)
}

func printSomething(value interface{}) {

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)

	}

}

func typeVal(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float: ", floatVal)
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String: ", stringVal)
	}
}

func add[T int | float64 | string](a, b T) T {
	// Generic function

	// Permite somar 2 valores independente do tipo
	return a + b

}

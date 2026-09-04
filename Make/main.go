package main

import "fmt"

type floatMap map[string]float64

func main() {

	// makingSlices()
	makingMap()

}

func (m floatMap) output() {
	fmt.Println(m)
}

func makingMap() {
	coursesRatings := make(floatMap, 3)

	coursesRatings["go"] = 4.7 // Sem o make, realoca memória sempre que define um novo map
	coursesRatings["React"] = 4.8
	coursesRatings["Angular"] = 4.8
	coursesRatings["C"] = 4.8 // Como passou de 3, esse realoca a memoria

	coursesRatings.output()

	// FOR LOOP
	for key, value := range coursesRatings {
		fmt.Println(key, value)
	}

}

func makingSlices() {
	userNames := make([]string, 2, 5) // Cria um vetor com 2 elementos (null)

	userNames[0] = "Max"
	userNames[1] = "Troy"
	userNames = append(userNames, "Steve")

	fmt.Println(userNames)

	// Demonstracao
	numeros := make([]int, 3, 5)
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros = append(numeros, 40)
	numeros = append(numeros, 50)

	fmt.Println(numeros)      // [10 20 30]
	fmt.Println(len(numeros)) // 3
	fmt.Println(cap(numeros)) // 5

	numeros = append(numeros, 60)
	fmt.Println(numeros)      // [0 0 0 40 50 60]
	fmt.Println(len(numeros)) // 6
	fmt.Println(cap(numeros)) // capacidade maior que 5

	// FOR LOOP
	for index, value := range userNames {
		fmt.Println(index, value)
	}

}

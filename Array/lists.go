package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {

	// basicoArraySlice()
	dynamicArray()
}

func basicoArraySlice() {

	var productNames [4]string = [4]string{"book"}
	prices := [4]float64{1.1, 2.2, 3.3, 4.4}

	productNames[2] = "Carpet"

	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[2])

	// Slice

	fmt.Println("\nBásico de slice")
	fmt.Println(prices[1:3])
	fmt.Println(prices[:3])
	fmt.Println(prices[1:])
	fmt.Println(prices[1:4])

	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]
	fmt.Print(highlightedPrices)

	// ---------------------------------------------

	fmt.Println("\n\nAprofundamento do slice")
	featuredPrices[0] = 99 // Slice é apenas uma parte da memória, logo muda o vetor original.
	fmt.Println(prices)
	// len: number of itens in a slice or array
	// cap:
	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
}

func dynamicArray() {

	prices := []float64{1.1, 2.2} // Omitir o tamanho

	updatedPrices := append(prices, 3.3, 4.4, 5.5) // Posso adicionar quantos eu quiser

	fmt.Println(updatedPrices, prices)

	// Remover o primeiro elemento (Não tem uma função específica)
	prices = prices[1:]
	fmt.Println(prices)

	// Merge 2 arrays
	discountPrices := []float64{10, 20}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

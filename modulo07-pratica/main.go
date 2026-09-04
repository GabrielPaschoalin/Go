package main

import "fmt"

type Product struct {
	title string
	price float64
	id    string
}

func newProduct(title string, price float64, id string) Product {
	return Product{
		title, price, id,
	}
}

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	array := [3]string{"|Tennis|", "|Watch soccer|", "|Programming|"}
	fmt.Println(array)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(array[0])
	fmt.Println(array[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	array1 := array[:2]
	array2 := array[0:2]
	fmt.Println(array1)
	fmt.Println(array2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	array3 := array1[1:3]
	fmt.Println(array3)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	array4 := []string{"|Learn Go|", "|Get a better job|"}
	fmt.Println(array4)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	array4[1] = "|Learn english|"
	array4 = append(array4, "|salary raise|")
	fmt.Println(array4)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	productArray := []Product{
		newProduct("Produto 1", 1.1, "001"), newProduct("Produto 2", 2.2, "002"),
	}
	productArray = append(productArray, newProduct("Produto 3", 3.3, "003"))

	fmt.Println(productArray)
}

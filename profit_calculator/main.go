package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err1 := printAndScan("Revenue: ")
	expenses, err2 := printAndScan("Expenses: ")
	taxRate, err3 := printAndScan("Tax rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	EBT, profit, ratio := calcularValores(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f \n", EBT)
	fmt.Printf("Profit: %.2f \n", profit)
	fmt.Printf("Ratio: %.2f \n", ratio)

	result := fmt.Sprintf("EBT: %.2f \nExpenses: %.2f \nTax rate: %.2f", EBT, profit, ratio)
	os.WriteFile("profit_calculator.txt", []byte(result), 0644)

}

func printAndScan(text string) (float64, error) {

	var output float64

	fmt.Print(text)
	fmt.Scan(&output)

	if output <= 0 {
		return 0, errors.New("Value must be a positive number!")
	}

	return output, nil
}

func calcularValores(revenue, expenses, taxRate float64) (EBT, profit, ratio float64) {

	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit

	return EBT, profit, ratio
}

func print(text string, value float64) {
	fmt.Print(text)
	fmt.Println(value)

}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

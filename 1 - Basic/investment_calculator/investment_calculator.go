package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {

	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years: ")
	fmt.Scan(&years)

	outputText("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := futureValues(investmentAmount, expectedReturnRate, years)

	// formattedFV := fmt.Sprintf("Future Value: %.2f \n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f \n", futureRealValue)
	// fmt.Print(formattedFV, formattedRFV)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value (adjusted for Inflation): ", futureRealValue)

	fmt.Printf("Future Value: %.2f \n", futureValue)
	fmt.Printf("Future Value (adjusted for Inflation): %.2f \n", futureRealValue)

	// fmt.Printf(`
	// Future Value: %.2f
	// Future Value (adjusted for Inflation): %.2f `,
	// 	futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func futureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {

	fv = investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	frv = fv / math.Pow(1+inflationRate/100, years)

	return fv, frv

	// Dá pra adicionar só o return return

}

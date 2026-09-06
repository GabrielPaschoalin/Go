package main

import (
	"fmt"
	filemanager "gabriel/calculator/fileManager"
	"gabriel/calculator/prices"
	"os"
	"strconv"
	"strings"
)

type mapFloat map[float64][]float64

func main() {

	// priceInput, err := []float64{12, 40, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, tax := range taxRates {
		fm := filemanager.New("inputPfdsfrice.txt", fmt.Sprintf("./files/outputPrice_%.0f.json", tax*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.New(tax, fm)
		err := priceJob.Process()

		if err != nil {
			fmt.Println(err)
			break
		}
	}

}

func ReadFile() ([]float64, error) {

	// Essa foi a que eu fiz (funcionou)
	content, err := os.ReadFile("inputPrice.txt")

	if err != nil {
		return nil, err
	}

	valueText := string(content)

	vec := strings.Fields(valueText)

	var array []float64 = make([]float64, len(vec))
	var floatValue float64

	for index, value := range vec {
		floatValue, err = strconv.ParseFloat(value, 64)

		array[index] = floatValue
	}

	if err != nil {
		return nil, err
	}

	return array, nil
}

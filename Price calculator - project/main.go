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
	doneChans := make([]chan bool, len(taxRates))
	errorChan := make([]chan error, len(taxRates))

	for index, tax := range taxRates {

		doneChans[index] = make(chan bool)
		errorChan[index] = make(chan error)

		fm := filemanager.New("inputPrice.txt", fmt.Sprintf("./files/outputPrice_%.0f.json", tax*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.New(tax, fm)
		go priceJob.Process(doneChans[index], errorChan[index]) // Não retorna valor

	}

	for index := range taxRates {
		select {
		case err := <-errorChan[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")

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

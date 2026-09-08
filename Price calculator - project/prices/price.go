package prices

import (
	"encoding/json"
	"fmt"
	iomanager "gabriel/calculator/IOManager"
	"gabriel/calculator/conversion"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

// IOManager         filemanager.FileManager `json:"-"` // Don't show the field

func New(taxRate float64, iom iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: iom,
	}
}

func (t *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {

	err := t.LoadData()

	if err != nil {
		errorChan <- err
		return
	}

	pricesPostTax := make(map[string]string, len(t.InputPrices))

	for _, price := range t.InputPrices {
		taxIncludedPrice := price * (1 + t.TaxRate)
		pricesPostTax[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(pricesPostTax)

	t.TaxIncludedPrices = pricesPostTax

	t.IOManager.WriteResult(t)

	doneChan <- true
}

func (t *TaxIncludedPriceJob) LoadData() error {

	lines, err := t.IOManager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.ConvertStringArrayToFloat(lines)

	if err != nil {
		return err
	}

	t.InputPrices = prices
	return nil
}

func (t TaxIncludedPriceJob) WriteData() {

	// CRIADO POR MIM - FUNCIONAL
	jsonData, err := json.Marshal(t.TaxIncludedPrices)

	if err != nil {
		fmt.Println("Error on JSON conversion")
		return
	}

	err = os.WriteFile("outputPrice.json", jsonData, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

}

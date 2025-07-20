package prices

import (
	"fmt"

	"example.com/price-caculator/conversion"
	"example.com/price-caculator/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	Prices            []float64           `json:"prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) Calculate() error {
	err := job.LoadData()

	if err != nil {
		fmt.Println("Error loading data:", err)
		return err
	}
	result := make(map[string]string)
	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result

	return job.IOManager.WriteJson(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println("Error reading lines from file:", err)
		return err
	}

	prices, err := conversion.StringsToFloat64s(lines)

	if err != nil {
		fmt.Println("Error converting strings to float64:", err)
		return err
	}

	job.Prices = prices
	return nil
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		Prices:    []float64{10, 20, 30},
		IOManager: iom,
	}
}

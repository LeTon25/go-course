package main

import (
	"fmt"

	"example.com/price-caculator/filemanager"
	"example.com/price-caculator/prices"
)

func main() {
	var taxRates []float64
	taxRates = []float64{0.5, 7, 15.2, 20.12}

	for _, taxRate := range taxRates {
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("tax_included_prices_%.0f.json", taxRate))
		err := prices.NewTaxIncludedPriceJob(fm, taxRate).Calculate()
		if err != nil {
			fmt.Printf("Error calculating tax included prices for tax rate %.2f: %v\n", taxRate, err)
			continue
		}
	}
}

package main

import (
	"fmt"

	"example.com/price-caculator/filemanager"
	"example.com/price-caculator/prices"
)

func main() {
	var taxRates []float64
	taxRates = []float64{0.5, 7, 15.2, 20.12}
	doneChans := make([]chan bool, len(taxRates))
	errChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errChans[index] = make(chan error)
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("tax_included_prices_%.0f.json", taxRate))
		go prices.NewTaxIncludedPriceJob(fm, taxRate).Calculate(doneChans[index], errChans[index])
	}

	for index, x := range taxRates {
		select {
		case err := <-errChans[index]:
			if err != nil {
				fmt.Printf("Error processing tax rate %.2f: %v\n", x, err)
			}
		case <-doneChans[index]:
			fmt.Printf("Successfully processed tax rate %.2f\n", x)
		}
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}

package main

import (
	"fmt"
	"math"
)

const INFLATION_RATE float64 = 2

func main() {
	run_bank()
}
func run_investment_caculator() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years : ")
	fmt.Scan(&years)

	result := caculateFutureValues(investmentAmount, expectedReturnRate, years)
	fmt.Printf("Future value : %.1f", result)
}
func outputText(text1 string) {
	println(text1)
}

func caculateFutureValues(investmentAmount, expectedReturnRate, years float64) float64 {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+INFLATION_RATE/100, years)

	return futureRealValue
}

package main

import (
	"fmt"
)

type TaxBracket struct {
	BracketLimit float64
	TaxRate      float64
}

func calculateTax(income float64, taxBrackets []TaxBracket) float64 {
	taxResult := 0.0
	remainingIncome := income

	for _, bracket := range taxBrackets {
		bracketLimit := bracket.BracketLimit
		taxRate := bracket.TaxRate
		if remainingIncome > 0 {
			taxableAmount := income
			if remainingIncome > bracketLimit {
				taxableAmount = bracketLimit
			}
			taxResult += taxableAmount * taxRate
			remainingIncome -= taxableAmount
		} else {
			break
		}
	}

	return taxResult
}

func main() {
	taxBracketsList := []TaxBracket{
		{30000000, 0.05},
		{100000000, 0.10},
		{999999999999, 0.20},
	}

	fmt.Print("input: ")
	var incomeInput float64
	_, err := fmt.Scanf("%f", &incomeInput)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	tax := calculateTax(incomeInput, taxBracketsList)
	formattedTax := fmt.Sprintf("%.0f", tax)

	fmt.Println("output:", formattedTax)
}

package main

import (
	"fmt"
)

func main() {
	const usdInEuro float64 = 0.86
	const usdInRub float64 = 72.76
	const euroInRub float64 = usdInRub / usdInEuro
	
	fmt.Println("Курс EURO -> RUB: ", euroInRub)
}

func getuserInPut() (float64, string, string) {
	var amount float64
	var fromCurrency string
	var toCurrency string

	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)
	fmt.Print("Из валюты: ")
	fmt.Scan(&fromCurrency)
	fmt.Print("В валюту: ")
	fmt.Scan(&toCurrency)
	return amount, fromCurrency, toCurrency
}

func convertCurrency(amount float64, fromCurrency string,toCurrency string) {

}

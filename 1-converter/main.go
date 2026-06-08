package main

import (
	"fmt"
)

const usdInEuro float64 = 0.86
const usdInRUB float64 = 72.76
const euroInRUB float64 = usdInRUB / usdInEuro

var currencyRates = map[string]float64{
	"USD": 1.0,
	"EUR": 1/usdInEuro,
	"RUB": 1/usdInRUB,
}

func main() {
	fmt.Println("___ Конвертер валют ___")
	fromCurrency := getCurrency("Введите исходную валюту")
	amount := getAmount()
	toCurrency := getCurrency("Введите целевую валюту")
	result := convertCurrency(amount, fromCurrency, toCurrency)
	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}

func getCurrency(message string) string {
	var currency string
	
	for {
		fmt.Println(message)
		fmt.Print("Доступные варианты: USD, EUR, RUB: ")
		_, err := fmt.Scan(&currency)
		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте еще раз.")
			continue
		}
		if currency == "USD" || currency == "EUR" || currency == "RUB" {
			return currency
			}
	fmt.Println(`Ошибка: такой валюты нет. 
Введите одну из трех USD / EUR / RUB.`)
	}
}

func getAmount() float64 {
	var amount float64
	
	for{
		fmt.Print("Введте сумму: ")
		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Ошибка: введите корректное число.")
			continue
		}
		
		if amount > 0 {
			return amount
		}
		 fmt.Println("Ошибка: сумма должна быть больше 0.")
	}
}	
	
func convertCurrency(amount float64, fromCurrency string,toCurrency string) float64 {
		if fromCurrency == toCurrency {
			return amount
		}
		fromRate := currencyRates[fromCurrency]
		toRate := currencyRates[toCurrency]
		return amount * fromRate / toRate
}
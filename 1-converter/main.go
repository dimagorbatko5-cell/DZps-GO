package main

import (
	"fmt"
	// "golang.org/x/text/message"
)

const usdInEuro float64 = 0.86
const usdInRUB float64 = 72.76
const euroInRUB float64 = usdInRUB / usdInEuro

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
		fmt.Scan(&currency)
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
		fmt.Scan(&amount)
		
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
		switch fromCurrency {
		case "USD":
			if toCurrency == "EUR" {
				return amount * usdInEuro
			}
			if toCurrency == "RUB" {
				return amount * usdInRUB
			}
			
		case "EUR":
			if toCurrency == "USD" {
				return amount / usdInEuro
			}
			if toCurrency == "RUB" {
				return amount * euroInRUB
			}
			
		case "RUB":
			if toCurrency == "USD" {
				return amount / usdInRUB
			}
			if toCurrency == "EUR" {
				return amount / euroInRUB
			}
		}
		return 0
}
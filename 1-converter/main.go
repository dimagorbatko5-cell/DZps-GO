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
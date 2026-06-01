package main

import "fmt"

func main() {

	
	transaktions := []float64{}
	for {
		transaktion := scanTransaction()
		if transaktion == 0 {
			break
	}
	transaktions = append(transaktions, transaktion)
	
	}	
	fmt.Println(transaktions)
}

func scanTransaction() float64 {
	var transaktion float64
	fmt.Print("Введите сумму транзакции или (n) для выхода: ")
	fmt.Scan(&transaktion)
	return transaktion
}
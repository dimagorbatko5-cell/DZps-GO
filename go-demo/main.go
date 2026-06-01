package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	defer func ()  {
		r := recover()
		if r != nil {
			fmt.Println("Recover ", r)
		}
 	}()
	fmt.Println("___ Калькулятор индекса массы тела ___" )
	for {
		userKg, userHaight := getuserInput()
		IMT, err := calculateIMT(userKg, userHaight)
		if err != nil {
			// fmt.Println(err)
			// continue
			panic(err)
		}
		outputResult(IMT)
		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}
	
}	

func outputResult(IMT float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.1f", IMT)
	fmt.Println(result)
	switch {
	case IMT < 16:
		fmt.Println("У Вас сильный дефицит массы тела")
	case IMT < 18.5:
		fmt.Println("У Вас дефицит массы тела")
	case IMT < 25:
		fmt.Println("У Вас норма массы тела")
	case IMT < 30:
		fmt.Println("У Вас избаточная масса тела")
	case IMT < 35:
		fmt.Println("У Вас 1-я степень ожирения")
	case IMT < 40:
		fmt.Println("У Вас 2-я степень ожирения")
	default:
		fmt.Println("У Вас 3-я степень ожирения")
	}
}

func calculateIMT (userKg float64, userHaight float64) (float64, error) {
	if userKg <= 0 || userHaight <=0 {
		return 0, errors.New("Не указан вес или высота, введите верное значение.")
	}
	const IMTPower = 2
	IMT := userKg / math.Pow(userHaight/100, IMTPower)
	return IMT, nil 
}

func getuserInput () (float64, float64) {
	var userHaight float64
	var userKg float64
	fmt.Print("Введите свой рост в см.: ")
	fmt.Scan(&userHaight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userKg, userHaight
}	

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("Хотите сделать еще один расчет? (Да/Нет)")
	fmt.Scan(&userChoise)
	if userChoise ==  "Да" || userChoise ==  "да" {
		return true
	}
	return false
}
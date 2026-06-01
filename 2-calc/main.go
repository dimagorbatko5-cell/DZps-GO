package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operation string
	var input string

	fmt.Print("Введите операцию SUM, AVG или MED: ")
	fmt.Scanln(&operation)

	fmt.Print("Введите числа через запятую: ")
	fmt.Scanln(&input)

	numbers, err := parseNumbers(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	result, err := calculate(operation, numbers)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}

func parseNumbers(input string) ([]float64, error) {
	parts := strings.Split(input, ",")

	numbers := make([]float64, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)

		number, err := strconv.ParseFloat(part, 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func calculate(operation string, numbers []float64) (float64, error) {
	switch operation {
	case "SUM":
		return sum(numbers), nil
	case "AVG":
		return avg(numbers), nil
	case "MED":
		return median(numbers), nil
	default:
		return 0, fmt.Errorf("неизвестная операция: %s", operation)
	}
}

func sum(numbers []float64) float64 {
	result := 0.0

	for _, number := range numbers {
		result += number
	}

	return result
}

func avg(numbers []float64) float64 {
	return sum(numbers) / float64(len(numbers))
}

func median(numbers []float64) float64 {
	sort.Float64s(numbers)

	length := len(numbers)
	middle := length / 2

	if length%2 == 1 {
		return numbers[middle]
	}

	return (numbers[middle-1] + numbers[middle]) / 2
}

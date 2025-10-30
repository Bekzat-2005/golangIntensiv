package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Поймана паника:", r)
		}
	}()

	var a, b float64
	var op string

	fmt.Print("Введите выражение (например: 10 / 2): ")
	fmt.Scan(&a, &op, &b)

	result, err := calculate(a, b, op)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат:", result)
}

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("на ноль делить нельзя")
		}
		return a / b, nil
	default:
		panic("неизвестная операция: " + op)
	}
}

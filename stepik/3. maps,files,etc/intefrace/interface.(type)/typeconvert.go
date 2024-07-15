package main

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	var v1, v2 bool
	var res1, res2, result float64
	switch value1.(type) {
	case float64:
		v1 = true
		res1 = value1.(float64)
	default:
		fmt.Printf("value=%v: %T", value1, value1)
		return
	}
	switch value2.(type) {
	case float64:
		v2 = true
		res2 = value2.(float64)
	default:
		fmt.Printf("value=%v: %T", value2, value2)
		return
	}
	if v1 && v2 {
		switch operation {
		case "+":
			result = res1 + res2
		case "-":
			result = res1 - res2
		case "*":
			result = res1 * res2
		case "/":
			result = res1 / res2
		default:
			fmt.Printf("Неизвестная операция")
			return
		}
	}
	fmt.Printf("%.4f", result)
}

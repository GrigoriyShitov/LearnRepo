package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"strings"
)

type Stringer interface {
	String() string
}
type Battery struct {
	charge string
}

func (b Battery) String() string {
	var result string
	var cnt int
	cnt = strings.Count(b.charge, "1")
	for i := 0; i < 10; i++ {
		if i >= 10-cnt {
			result += "X"
		} else {
			result += " "
		}

	}
	return fmt.Sprintf("[%s]", result)
}

func main() {
	var batteryForTest Battery
	fmt.Scan(&batteryForTest.charge)
	// batteryForTest - не забывайте об имени
} //Скобка, закрывающая функцию main() вам не видна, но она есть

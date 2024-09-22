package main

import "fmt"

func ExampleRune() {
	// Поступим также, как в работе с типом []byte
	rs := []rune("Это срез рун")

	// Итерируясь мы будем заменять символ 'р' на '*'
	for i := range rs {
		if rs[i] == 'р' {
			rs[i] = '*'
		}
	}
	fmt.Printf("Измененнный срез в виде строки: %s\n", string(rs))

	// Output:
	// Измененнный срез в виде строки: Это с*ез *ун
}

func main() {

	ExampleRune()
}

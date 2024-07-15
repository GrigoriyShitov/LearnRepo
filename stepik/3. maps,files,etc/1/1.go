package main

import "fmt"

func main() {

	m := make(map[int]int)
	var digit int
	for i := 0; i < 10; i++ {
		fmt.Scan(&digit)
		if value, InMap := m[digit]; InMap {
			fmt.Print(value, " ")
		} else {
			m[digit] = work(digit)
			fmt.Print(m[digit], " ")
		}
	}
}

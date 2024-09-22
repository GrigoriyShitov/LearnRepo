package main

import (
	"fmt"
)

func main() {
	var x string
	fmt.Scan(&x)
	result := []rune()
	for i, elem := range x {
		if i%2 != 0 {
			result += elem
		}
	}

}

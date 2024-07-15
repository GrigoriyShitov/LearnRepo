package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)
	result, _ := strconv.Atoi(string(str[0]))
	for i := range str {
		digit, _ := strconv.Atoi(string(str[i]))
		if result < digit {
			result = digit
		}
	}

	fmt.Println(result)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str, result string
	fmt.Scan(&str)
	for i := range str {
		digit, _ := strconv.Atoi(string(str[i]))
		digit *= digit
		fmt.Print(digit)
		result += string(digit)
	}

}

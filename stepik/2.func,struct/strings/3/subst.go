package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, s string
	fmt.Scan(&x, &s)
	index := strings.Index(x, s)
	if index != -1 {
		fmt.Println(index)
	} else {
		fmt.Println("-1")
	}
}

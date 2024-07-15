package main

import (
	"fmt"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	text := []rune(s)
	if len(text) >= 5 {
		for i := range text {
			if unicode.IsDigit(text[i]) || unicode.Is(unicode.Latin, text[i]) {
				continue
			} else {
				fmt.Println("Wrong password")
				return
			}
		}
	} else {
		fmt.Println("Wrong password")
		return
	}
	fmt.Println("Ok")
}

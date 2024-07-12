package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text := []rune(s)
	textlen := len(text) - 2
	// fmt.Print("pizda:  ", string(text[textlen]))
	if unicode.IsUpper(text[0]) && text[textlen] == '.' {
		fmt.Println("Right")

	} else {
		fmt.Println("Wrong")
	}
	//if (unicode.IsUpper(s[0]))
	//fmt.Printf("%v,\n %d", string(text), len(text))
}

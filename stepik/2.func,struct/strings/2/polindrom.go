package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text := []rune(s)
	textlen := len(text) - 2
	for i := 0; i <= textlen/2; i++ {
		if text[i] != text[textlen-i] {
			fmt.Printf("Нет")
			return
		}
	}
	fmt.Println("Палиндром")
}

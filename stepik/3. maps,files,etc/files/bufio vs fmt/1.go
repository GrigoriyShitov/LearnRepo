package main

import (
	"bufio"
	"os"
	"strconv"
)

// package main уже объявлен.
func main() {
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.
	scaner := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var sum int
	for scaner.Scan() {
		s := scaner.Text()
		digit, _ := strconv.Atoi(string(s))
		sum += digit
	}
	w.WriteString(strconv.Itoa(sum))
	w.Flush()
}

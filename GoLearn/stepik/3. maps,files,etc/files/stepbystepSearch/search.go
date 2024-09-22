package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("task.data")
	defer file.Close()
	var pos int
	if err != nil {
		fmt.Println("error")
	}
	rd := bufio.NewReader(file)
	for {
		s, err := rd.ReadString(';')
		digit, _ := strconv.Atoi(strings.Trim(s, ";"))
		fmt.Println(digit)
		if err == io.EOF || digit == 0 {
			pos++
			break
		} else {
			pos++
		}
	}
	fmt.Println(pos)
}

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func adding(s1 string, s2 string) int64 {
	var s1final, s2final string

	for _, elem := range s1 {
		if unicode.IsDigit(elem) {
			s1final += string(elem)
		}
	}

	for _, elem := range s2 {
		if unicode.IsDigit(elem) {
			s2final += string(elem)
		}
	}
	sum1, _ := strconv.Atoi(s1final)
	sum2, _ := strconv.Atoi(s2final)
	return int64(sum1 + sum2)

}

func main() {
	s1 := "%^80"
	s2 := "hhhhh20&&&&nd"
	fmt.Println(adding(s1, s2))
}

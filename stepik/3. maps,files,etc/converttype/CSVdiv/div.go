package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	s1, s2 := strings.Split(string(s), ";")[0], strings.Split(string(s), ";")[1]
	s1, s2 = strings.Replace(s1, ",", ".", 1), strings.Replace(s2, ",", ".", 1)
	s1, s2 = strings.Replace(s1, " ", "", 1), strings.Replace(s2, " ", "", 1)
	//fmt.Println(s1, s2)
	d1, _ := strconv.ParseFloat(s1, 64)
	d2, _ := strconv.ParseFloat(s2, 64)
	//fmt.Println(d1, d2)
	result := d1 / d2
	fmt.Printf("%.4f", result)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	date1, _ := buf.ReadString(',')
	date1 = strings.Trim(date1, ",")
	t1, err := time.Parse("02.01.2006 15:04:05", date1)
	if err != nil {
		fmt.Println(err)
	}
	date2, _ := buf.ReadString('\n')
	date2 = strings.Trim(date2, "\n")
	t2, err := time.Parse("02.01.2006 15:04:05", date2)
	if err != nil {
		fmt.Println(err)
	}
	if t1.After(t2) {
		fmt.Println(t1.Sub(t2))
	} else {
		fmt.Println(t2.Sub(t1))
	}

}

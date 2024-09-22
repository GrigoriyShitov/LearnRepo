package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	buf = strings.Trim(buf, "\n")
	t, err := time.Parse("2006-01-02 15:04:05", buf)
	if err != nil {
		fmt.Println(err)
	}
	if t.Hour() > 12 {
		t = t.AddDate(0, 0, 1)
	}
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}

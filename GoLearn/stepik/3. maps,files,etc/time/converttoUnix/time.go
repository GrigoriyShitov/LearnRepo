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
	t, err := time.Parse(time.RFC3339, buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Format(time.UnixDate))
}

package main

import (
	"fmt"
)

func work() {
	fmt.Println("")
}

func main() {
	done := make(chan struct{})
	go func() {
		work()
		close(done)
	}()
	<-done
}

package main

import "fmt"

func task(c chan int, n int) {
	n++
	c <- n
}

func main() {
	c := make(chan int, 10)
	task(c, 1)
	fmt.Println(<-c)
}

package main

import "fmt"

func task2(c chan string, s string) {
	s = s + " "
	for i := 0; i < 5; i++ {
		c <- s
	}

}

func main() {
	c := make(chan string, 5)
	str := "hello"

	go task2(c, str)

	for i := 0; i < 5; i++ {
		fmt.Print(<-c)
	}

}

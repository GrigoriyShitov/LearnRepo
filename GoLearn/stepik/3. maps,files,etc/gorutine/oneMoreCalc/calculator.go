package main

import "fmt"

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	output := make(chan int)

	go func(output chan int) {
		defer close(output)
		var sum int
		for {
			select {
			case <-done:
				output <- sum
				return
			case x := <-arguments:
				sum += x

			}
		}

	}(output)
	return output
}

func main() {
	arguments := make(chan int)
	done := make(chan struct{})
	result := calculator(arguments, done)
	for i := 0; i < 10; i++ {
		arguments <- i
	}
	close(done)
	fmt.Println(<-result)
}

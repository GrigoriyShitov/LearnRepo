package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)

	go func(output chan int) {
		defer close(output)
		select {
		case x := <-firstChan:
			output <- x * x
		case x := <-secondChan:
			output <- x * 3
		case <-stopChan:
			break
		}
	}(output)
	return output
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})
	r := calculator(ch1, ch2, stop)
	close(stop)
	//ch1 <- 3
	ch2 <- 5
	fmt.Println(<-r)

}

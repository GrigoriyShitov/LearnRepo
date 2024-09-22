package main

import "fmt"

func removeDuplicates(input, output chan string) {
	var prev string
	for i := range input {
		if i != prev {
			output <- i
		}
		prev = i
	}
	close(output)
}

func main() {
	input := make(chan string, 100)
	output := make(chan string, 100)
	go removeDuplicates(input, output)
	go func() {
		defer close(input)

		for _, r := range "112334456" {
			input <- string(r)
		}
	}()

	for x := range output {
		fmt.Print(x)
	}

}

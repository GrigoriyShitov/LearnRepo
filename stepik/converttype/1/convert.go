package main

import "fmt"

func convert(d uint64) uint16 {
	var dLess uint16
	dLess = uint16(d)
	return dLess

}

func main() {
	d := 64
	fmt.Println(convert(uint64(d)))
}

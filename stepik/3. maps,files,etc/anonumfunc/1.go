package main

import (
	"fmt"
	"strconv"
)

func main() {

	fn := func(i uint) uint {
		str := strconv.Itoa(int(i))
		var result string
		for _, elem := range str {
			digit, _ := strconv.Atoi(string(elem))
			if digit%2 == 0 && digit != 0 {
				result += string(elem)
			}
		}
		//fmt.Println(string(result))
		returnValue, _ := strconv.Atoi(result)
		if returnValue == 0 {
			return uint(100)
		}

		return uint(returnValue)
	}
	fmt.Println(fn(0))
}

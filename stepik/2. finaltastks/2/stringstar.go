package main

import "fmt"

func main() {
	var str, result string
	fmt.Scan(&str)
	for i := range str {
		result += string(str[i])
		if i == len(str)-1 {
			break
		}
		result += "*"
	}

	fmt.Println(result)
}

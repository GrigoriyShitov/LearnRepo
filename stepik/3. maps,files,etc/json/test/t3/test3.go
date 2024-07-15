package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

func main() {

	data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)
	var s myStruct

	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v", s) // {John Connor 35 true [15 11 37]}
}

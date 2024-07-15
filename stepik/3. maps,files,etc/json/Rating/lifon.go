package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type myStruct struct {
	ID       int
	Number   string
	Year     int
	Students []students
}
type students struct {
	LastName   string `json:"LastName"`
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}
type output struct {
	Average float64
}

func main() {
	file, _ := os.Open("test.json")
	data, _ := io.ReadAll(file)
	var s myStruct
	var out output
	var cntR, cntS int
	var avg float64
	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}
	for i := range s.Students {
		cntS++
		cntR += len(s.Students[i].Rating)
	}

	avg = float64(cntR) / float64(cntS)
	out.Average = avg
	data, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data)
	//fmt.Printf("%v", s.Students[1].Rating) // {John Connor 35 true [15 11 37]}
}

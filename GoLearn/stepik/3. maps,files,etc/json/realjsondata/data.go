package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type dataStruct struct {
	GlobalID       int    `json:"global_id"`
	SystemObjectID string `json:"system_object_id"`
	SignatureDate  string `json:"signature_date"`
	Razdel         string `json:"Razdel"`
	Kod            string `json:"Kod"`
	Name           string `json:"Name"`
	Idx            string `json:"Idx"`
}

type rawData struct {
	units []dataStruct
}

func main() {
	file, _ := os.Open("data-20190514T0100.json")
	data, _ := io.ReadAll(file)
	var result uint64
	var r rawData
	if err := json.Unmarshal(data, &r.units); err != nil {
		fmt.Println(err)
		return
	}
	for i := range r.units {
		result += uint64(r.units[i].GlobalID)
	}
	fmt.Println(result)
}

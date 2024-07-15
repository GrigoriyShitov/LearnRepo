package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func walkFunc(path string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if strings.Contains(info.Name(), ".txt") {
		file, err := os.Open(path)
		if err != nil {
			err := fmt.Errorf("ошибка открытия файла %s", info.Name())
			return err
		}

		defer file.Close()
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			err := fmt.Errorf("ошибки при считывании файла =(%s", err)
			return err
		}

		if len(records) == 10 {
			fmt.Println(records[4][2])
		}

		return nil
	}

	return nil

}

func main() {
	const root = "../task"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

}

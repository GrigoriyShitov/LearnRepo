package main

import "fmt"

func main() {

	groupCity := map[int][]string{
		10:   []string{"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Станица"},        // города с населением 100-999 тыс. человек
		1000: []string{"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	}
	inMap := false
	for key, _ := range cityPopulation {

		for _, valueGR := range groupCity[100] {
			if valueGR == key {
				inMap = true
				break
			}
		}
		if !inMap {
			delete(cityPopulation, key)
		}
		inMap = false
	}
	x := func(fn func(i int) int, i int) func(int) int { return fn }(func(i int) int { return i + 1 }, 5)

	fmt.Println(cityPopulation)
	fmt.Printf("%T", x)
}

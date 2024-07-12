package main

import (
	"fmt"
	"math"
)

const p = 6.0
const v = 6.0
const k = 1296.0

func main() {

	fmt.Println(T())
}

func T() float64 {

	w := W()
	return 6 / w
}

func W() float64 {

	m := M()
	return math.Sqrt(k / m)
}

func M() float64 {
	return p * v
}

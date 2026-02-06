package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Cat struct {
	color string
}

func (c *Cat) changeColor() {
	c.color = "pink"
}

var (
	errTimeout error = errors.New("Timeout")
)

func main() {
	start := time.Now()

	ctx1, _ := context.WithTimeout(context.Background(), 3*time.Second)
	ctx2, _ := context.WithTimeout(ctx1, 2*time.Second)
	ctx3, _ := context.WithTimeout(ctx2, 5*time.Second)

	<-ctx3.Done()

	fmt.Println(time.Since(start).Round(time.Second).Seconds())
}

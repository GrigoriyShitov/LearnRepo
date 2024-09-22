package main

import (
	"fmt"
	"time"
)

const now = 1589570165

func main() {
	var min, sec int64
	fmt.Scanf("%d мин. %d сек.", &min, &sec)
	t := time.Unix(now+min*60+sec, 0).UTC()
	fmt.Println(t.Format(time.UnixDate))
}

package main

import (
	"fmt"

	"l_1/l1_24/point"
)

func main() {
	p1 := point.NewPoint(0, 0)
	p2 := point.NewPoint(1, 1)

	fmt.Println("Distance:", p1.Distance(p2))
}

package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)

	for _, t := range temps {
		var key int
		if t >= 0 { //positive - round to MIN
			key = int(math.Floor(t/10) * 10)
		} else { //negative - round to MAX
			key = int(math.Ceil(t/10) * 10)
		}
		groups[key] = append(groups[key], t)
	}

	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, v)
	}
}

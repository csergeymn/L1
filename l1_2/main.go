package main

import (
	"fmt"
)

type Result struct {
	Index int
	Value int
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	squares := make(chan Result, len(numbers))

	for i, n := range numbers {
		go func(idx, num int) {
			fmt.Println(num)
			squares <- Result{
				Index: idx,
				Value: num * num,
			}
		}(i, n)
	}

	output := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		res := <-squares
		output[res.Index] = res.Value
	}

	for i, val := range output {
		fmt.Printf("%d^2 = %d\n", numbers[i], val)
	}
}

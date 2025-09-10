package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)
	results := make(chan int)

	go func() {
		arr := []int{1, 5, 10, 15, 50}
		for _, v := range arr {
			numbers <- v
		}
		close(numbers)
	}()

	go func() {
		for x := range numbers {
			results <- x * 2
		}
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("✅ Pipeline finished")
}

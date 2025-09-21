package main

import (
	"fmt"
)

func quickSortAsc(arr []int) []int {
	fmt.Println("quickSort called")

	if len(arr) < 2 {
		return arr
	}

	medium := arr[len(arr)/2]

	var less []int
	var equal []int
	var greater []int

	for _, v := range arr {
		switch {
		case v < medium:

			less = append(less, v)
		case v == medium:
			equal = append(equal, v)
		case v > medium:
			greater = append(greater, v)
		}
	}

	return append(append(quickSortAsc(less), equal...), quickSortAsc(greater)...)
}

func main() {
	arr := []int{10, 5, 2, 3, 7, 6, 8, 1, 9, 4}
	fmt.Println("Original:", arr)

	sorted := quickSortAsc(arr)
	fmt.Println("Sorted:", sorted)
}

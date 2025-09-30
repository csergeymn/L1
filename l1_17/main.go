package main

import (
	"fmt"
)

func quickSortAsc(arr []int) []int {
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

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func main() {
	arr := []int{10, 20, 30, 50, 100, 1, 3, 2, 4, 8}
	arr = quickSortAsc(arr)
	fmt.Println("Array:", arr)

	testValues := []int{2, 4, 1, 100, 50, -1, 101, 33}

	for _, v := range testValues {
		res := "Not found"
		i := binarySearch(arr, v)
		if i > -1 {
			res = "Found"
		}
		fmt.Printf("Search %d -> index:%d: %v \n", v, i, res)
	}
}

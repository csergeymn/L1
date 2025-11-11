package main

import (
	"errors"
	"fmt"
)

func remove[T any](s []T, i int) ([]T, error) {
	if i < 0 || i >= len(s) {
		errStr := fmt.Sprintf("Index out of range: [%d]", i)
		return s, errors.New(errStr)
	}
	fmt.Printf("Deleting at index [%d]\n", i)
	copy(s[i:], s[i+1:])
	s[len(s)-1] = *new(T)
	return s[:len(s)-1], nil
}

func main() {
	arr := []string{"cat", "dog", "fish", "bird"}
	fmt.Println("Before:", arr)

	arr, err := remove(arr, 1)
	if err != nil {
		fmt.Printf("Error: %v for slice %v\n", err, arr)
	} else {
		fmt.Println("After:", arr)
	}

	arrInt := []int{10, 20, 30, 40, 50}
	fmt.Println("Before:", arrInt)

	arrInt, err = remove(arrInt, 10)
	if err != nil {
		fmt.Printf("Error: %v for slice %v\n", err, arrInt)
	} else {
		fmt.Println("After:", arrInt)
	}

	arrInt, err = remove(arrInt, 0)
	if err != nil {
		fmt.Printf("Error: %v\n for slice %v", err, arrInt)
	} else {
		fmt.Println("After:", arrInt)
	}
}

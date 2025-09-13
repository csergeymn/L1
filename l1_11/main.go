package main

import "fmt"

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	result := []int{}

	intersectMap := make(map[int]bool)
	for _, v := range A {
		intersectMap[v] = true
	}

	for _, v := range B {
		if intersectMap[v] {
			result = append(result, v)
			delete(intersectMap, v)
		}
	}

	fmt.Println("Intersection:", result)
}

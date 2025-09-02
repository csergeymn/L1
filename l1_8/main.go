package main

import (
	"fmt"
)

func setBit(n int64, i uint, bitValue int) int64 {
	if bitValue == 1 {
		// set i-th bit to 1
		n = n | (1 << i)
	} else {
		// set i-th bit to 0
		n = n &^ (1 << i) // &^ = AND NOT
	}
	return n
}

func main() {
	var num int64 = 15 // binary: 0101
	fmt.Printf("Original number: %d (binary %04b)\n", num, num)

	num = setBit(num, 1, 0)
	fmt.Printf("After setting 1st bit to 0: %d (binary %04b)\n", num, num)

	num = setBit(num, 2, 1)
	fmt.Printf("After setting 2nd bit to 1: %d (binary %04b)\n", num, num)
}

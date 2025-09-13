package main

import "fmt"

func main() {
	a, b := 500, 9999
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("After swap:  a=%d, b=%d\n", a, b)
}

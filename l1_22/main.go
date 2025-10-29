package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)

	a.SetString("1111111111111111111111111111111111111111", 10)
	b.SetString("2222222222222222222222222222222222222222", 10)

	fmt.Printf("a = %v\nb = %v\n\n", a, b)

	res := new(big.Int).Add(a, b)
	fmt.Printf("a + b = %v\n", res)
	res = new(big.Int).Sub(a, b)
	fmt.Printf("a - b = %v\n", res)
	res = new(big.Int).Mul(a, b)
	fmt.Printf("a * b = %v\n", res)
	res = new(big.Int).Div(b, a)
	fmt.Printf("b / a = %v\n", res)
}

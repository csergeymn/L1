package main

import (
	"fmt"
	"unicode"
)

func CheckUniqueSymbols(s string) bool {
	m := make(map[rune]struct{})
	for _, v := range s {
		v = unicode.ToLower(v)
		if _, has := m[v]; has {
			return false
		}

		m[v] = struct{}{}
	}

	return true
}

func main() {
	check := "Привет🙂 🙂"
	res := CheckUniqueSymbols(check)
	fmt.Println(res)
}

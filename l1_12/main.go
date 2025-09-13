package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, w := range words {
		set[w] = struct{}{}
	}

	var setWords []string
	for key := range set {
		setWords = append(setWords, key)
	}

	fmt.Printf("%v", setWords)

}

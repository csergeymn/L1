package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseRunes(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	runes := []rune(s)
	fmt.Printf("%v", runes)
	reverseRunes(runes, 0, len(runes)-1)
	fmt.Printf("%v \n", runes)
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverseRunes(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1]
	fmt.Println("Reverse words:", reverseWords(input))
}

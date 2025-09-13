package main

import "fmt"

func detectType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Printf("Type: int, Value: %v\n", v)
	case string:
		fmt.Printf("Type: string, Value: %v\n", v)
	case bool:
		fmt.Printf("Type: bool, Value: %v\n", v)
	case chan any:
		fmt.Printf("Type: chan, Value: %v\n", v)
	default:
		fmt.Printf("Unknown type, Value: %v\n", v)
	}
}

func main() {
	detectType(42)      // int
	detectType("hello") // string
	detectType(true)    // bool

	chInt := make(chan int)
	detectType(chInt) // chan

	chStr := make(chan string)
	detectType(chStr) // chan

	detectType(3.14) // unknown - no case branch
}

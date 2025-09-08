package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const bitSize = 64 // int64 has 64 bits

// setBit sets the i-th bit (1-based index) of n to bitValue (0 or 1).
func setBit(n int64, i uint, bitValue int) int64 {
	i-- // convert to 0-based index
	if bitValue == 1 {
		n = n | (1 << i) // set bit
	} else {
		n = n &^ (1 << i) // clear bit
	}
	return n
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nEnter a number (or 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			fmt.Println("👋 Exiting program...")
			break
		}

		num, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("❌ Invalid number")
			continue
		}

		fmt.Print("Enter bit position (1-64): ")
		if !scanner.Scan() {
			break
		}
		posStr := strings.TrimSpace(scanner.Text())
		pos, err := strconv.Atoi(posStr)
		if err != nil || pos <= 0 || pos > bitSize {
			fmt.Println("❌ Invalid bit position (must be between 1 and 64)")
			continue
		}

		fmt.Print("Enter bit value (0 or 1): ")
		if !scanner.Scan() {
			break
		}
		valStr := strings.TrimSpace(scanner.Text())
		val, err := strconv.Atoi(valStr)
		if err != nil || (val != 0 && val != 1) {
			fmt.Println("❌ Invalid bit value (must be 0 or 1)")
			continue
		}

		newNum := setBit(num, uint(pos), val)

		fmt.Printf("Original number: %d (binary %064b)\n", num, uint64(num))
		fmt.Printf("Modified number: %d (binary %064b)\n", newNum, uint64(newNum))
	}
}

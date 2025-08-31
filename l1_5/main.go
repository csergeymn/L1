package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter work time for program (maximum 30): ")
	var N int
	_, err := fmt.Fscan(reader, &N)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil || N <= 0 {
		log.Fatalln("Invalid work time for program")
		return
	}

	if N > 30 {
		log.Fatalln("It's too long for sample")
		return
	}

	ch := make(chan int)

	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(500 * time.Millisecond) // имитация работы
		}
	}()

	timeout := time.After(time.Duration(N) * time.Second)

	for {
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		case <-timeout:
			fmt.Println("⏱ Time's up! Exiting...")
			return
		}
	}
}

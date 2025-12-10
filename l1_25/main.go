package main

import (
	"fmt"
	"time"
)

func CustomSleep(d time.Duration) {
	timer := time.NewTimer(d)
	fmt.Printf("Timer started at: %v\n", time.Now())
	val := <-timer.C
	fmt.Printf("Timer channel got value: %v\n", val)
}

func main() {
	fmt.Println("Start")
	CustomSleep(2 * time.Second)
	fmt.Println("End")
}

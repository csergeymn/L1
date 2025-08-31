package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	stopFlag := false
	go func() {
		for !stopFlag {
			fmt.Println("[By condition] Working...")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("[By condition] goroutine stopped")
	}()

	doneChan := make(chan bool)
	go func() {
		for {
			select {
			case <-doneChan:
				fmt.Println("[Channel] goroutine stopped")
				return
			default:
				fmt.Println("[Channel] Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("[Context] goroutine stopped")
				return
			default:
				fmt.Println("[Context] Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("[Goexit] Working...")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("[Goexit] goroutine stopping by command immediately >>>")
		runtime.Goexit()
	}()

	stopTimer := time.After(2 * time.Second)
	go func() {
		for {
			select {
			case <-stopTimer:
				fmt.Println("[Timer] goroutine stopped")
				return
			default:
				fmt.Println("[Timer] Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(3 * time.Second)

	stopFlag = true
	doneChan <- true
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("All goroutines stopped")
}

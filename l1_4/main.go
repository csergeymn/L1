package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func worker(id int, jobs <-chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping...\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: jobs channel closed\n", id)
				return
			}
			fmt.Printf("Worker %d processed job %d\n", id, job)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of workers: ")
	var numWorkers int
	_, err := fmt.Fscan(reader, &numWorkers)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil || numWorkers <= 0 {
		fmt.Println("Invalid number of workers")
		return
	}

	jobs := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, ctx)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		<-sigChan
		fmt.Println("\nReceived interrupt, stopping...")
		cancel()
		close(jobs)
	}()

	jobID := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Main stopping...")
			return
		default:
			jobs <- jobID
			fmt.Printf("Main: sent job %d\n", jobID)
			jobID++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

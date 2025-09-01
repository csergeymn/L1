package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d processed job %d\n", id, job)
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

	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs)
	}

	jobID := 100
	for {
		jobs <- jobID
		fmt.Printf("Main: sent job %d\n", jobID)
		jobID++
		time.Sleep(500 * time.Millisecond)
	}
}

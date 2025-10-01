package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	i  int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.i++
	c.mu.Unlock()
}

func (c *Counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.i
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		fmt.Printf("Goroutine started № %d\n", i)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.getValue())
}

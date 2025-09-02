package main

import (
	"fmt"
	"sync"
)

var (
	concurrentMap   = make(map[string]int)
	concurrentMapMu sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Goroutine %d started\r\n", i)
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			concurrentMapMu.Lock()
			defer concurrentMapMu.Unlock()
			concurrentMap[key] = i
		}(i)
	}

	wg.Wait()

	fmt.Println("Result:", concurrentMap)
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		fmt.Println(key, concurrentMap[key])
	}
}

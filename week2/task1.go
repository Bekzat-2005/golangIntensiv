package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		mu.Lock()
		counter++
		mu.Unlock()
		results <- job * 2
		fmt.Printf("Worker %d processed job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}

	fmt.Println("Total jobs processed:", counter)
}

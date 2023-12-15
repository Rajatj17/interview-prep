package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(jobs <-chan int, wSeq int, wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Println("From Worker", wSeq, "Job", job)
		time.Sleep(time.Second)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 100)

	wg.Add(10)
	for i := 0; i < 2; i++ {
		go worker(jobs, i, &wg)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
}

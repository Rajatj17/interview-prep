package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

func newFunc() {
	counter = counter + 1
}

func main() {
	var once sync.Once

	for i := 0; i < 100; i++ {
		go func() {
			once.Do(newFunc)
		}()
	}

	time.Sleep(4 * time.Second)
	fmt.Printf("Count is %d\n", counter)
}

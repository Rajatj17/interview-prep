package main

import (
	"fmt"
	"sync"
)

var done = false

func add(a, b int, wg *sync.WaitGroup) int {
	defer wg.Done()
	sum := a + b
	fmt.Println(sum)
	return sum
}

func write(c *sync.Cond) {

}

func read(c *sync.Cond) {
	c.L.Lock()

	for !done {
		c.Wait()
	}

	fmt.Println("Read")
	c.L.Unlock()
}

func main() {
	a, b := 1, 2

	var c int
	var wg sync.WaitGroup
	var m sync.Mutex
	cond := sync.NewCond(&m)

	wg.Add(1)
	go add(a, b, &wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.Lock()
		c += 1
		m.Unlock()
	}()

	go read(cond)
	go read(cond)
	go write(cond)
	wg.Wait()

	fmt.Println("done", c)
}

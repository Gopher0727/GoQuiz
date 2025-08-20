package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	cnt int
	mu  sync.RWMutex
}

func (c *Counter) get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cnt
}

func (c *Counter) add(x int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cnt += x
}

func main() {
	var wg sync.WaitGroup

	siz := 100
	wg.Add(siz)

	var c Counter
	for i := range siz {
		go func() {
			defer wg.Done()

			c.get()
			c.add(i)
		}()
	}

	wg.Wait()

	fmt.Printf("Done! x = %v\n", c.cnt)
}

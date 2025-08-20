package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once

	siz := 10
	wg.Add(siz)

	for range siz {
		go func() {
			defer wg.Done()

			once.Do(func() { fmt.Println("Hello") })
		}()
	}

	wg.Wait()

	fmt.Println("Done!")
}

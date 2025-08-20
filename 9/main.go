package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%d: Start...\n", id)
	time.Sleep(time.Second)
}

func main() {
	var wg sync.WaitGroup

	siz := 10
	wg.Add(10)

	for i := range siz {
		go work(i, &wg)
	}

	wg.Wait()

	fmt.Println("Done!")
}

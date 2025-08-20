package main

import (
	"fmt"
	"sync"
)

func write(ch chan<- int, value int) {
	ch <- value
}

func read(ch <-chan int) int {
	return <-ch
}

func main() {
	ch := make(chan int, 10)

	wg := sync.WaitGroup{}
	wg.Add(1)

	var x int
	go write(ch, 42)
	go func() {
		defer wg.Done()
		x = read(ch)
	}()

	wg.Wait()
	fmt.Println(x)
}

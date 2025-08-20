package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := range 10 {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()

	for i := range ch {
		fmt.Println(i)
	}
}

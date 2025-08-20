package main

import "fmt"

func main() {
	ch := make(chan any, 2)

	ch <- "abc"
	ch <- 10
	close(ch)

	for i := range ch {
		fmt.Println(i)
	}
}

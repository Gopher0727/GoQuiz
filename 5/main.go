package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Hello, world! This is a goroutine.")
	}()

	fmt.Println("Hello, world! This is the main function. Love from main function.")
}

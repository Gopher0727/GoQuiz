package main

import "fmt"

func get() (res int) {
	defer func() {
		res *= 10
	}()
	return 1
}

func main() {
	i := 1

	defer fmt.Println(i)
	defer fmt.Println(get())

	i++
}

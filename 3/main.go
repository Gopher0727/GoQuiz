package main

import "fmt"

func add(a []int) {
	a = append(a, 4)
}

func main() {
	a := []int{1, 2, 3}
	add(a)
	fmt.Println(a) // [1 2 3]
}

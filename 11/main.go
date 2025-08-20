package main

import (
	"fmt"
	"sync"
	"time"
)

type Singleton struct {
	x int
}

var instance *Singleton

var once sync.Once

func get() *Singleton {
	once.Do(
		func() {
			fmt.Println("Start...")
			time.Sleep(time.Second)
			instance = &Singleton{}
			instance.x += 100
		},
	)
	fmt.Println(instance.x)
	return instance
}

func main() {
	var wg sync.WaitGroup

	siz := 10
	wg.Add(siz)

	for range siz {
		go func() {
			defer wg.Done()

			instance = get()
		}()
	}

	wg.Wait()

	fmt.Println("Done!")
}

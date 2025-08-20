package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	cond := sync.NewCond(&mu)

	siz := 10
	wg.Add(siz)

	for i := range siz {
		go func(i int) {
			defer wg.Done()

			cond.L.Lock()
			fmt.Printf("%v ready!", i)
			cond.Wait() // Wait() 先解锁 cond.L，再阻塞当前 goroutine
			fmt.Printf("%v done!", i)
			cond.L.Unlock()
		}(i)
	}

	time.Sleep(3 * time.Second)

	go func() {
		defer wg.Done()
		// 需要保证目标 goroutine 阻塞，否则死锁，Signal() 同理
		cond.Broadcast()
	}()

	wg.Wait()
}

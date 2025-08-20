package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestCounter(t *testing.T) {
	// var counter int32     // 计数值
	var data atomic.Int32
	var wg sync.WaitGroup // 等待 goroutine 执行完毕

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				// counter++ // 非原子操作，存在竞争条件
				data.Add(1)
			}
		}()
	}
	wg.Wait()
	// fmt.Println(counter)
	fmt.Println(data.Load()) // 使用原子操作获取计数值
}

func main() {
	TestCounter(&testing.T{})
}

package main

import (
	"fmt"
	"runtime"
	"time"
)

// Coroutine（協程）
// 輕量級線程
// 非搶佔式多任務處理 由協程主動交出控制權
// 編譯器 / 解釋器 / 虛擬機層面的多任務
// 多個協程可能在一個或多個線程上運行
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // race condition
			for {
				// fmt.Printf("Hello from goroutine %d\n", i)
				a[i]++
				runtime.Gosched() //手動交出控制權
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

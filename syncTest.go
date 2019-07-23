package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建条件变量cond
	//condition := false
	c := sync.NewCond(&sync.Mutex{})

	// 创建线程
	go func() {
		time.Sleep(time.Second * 5)
		c.Signal()
	}()

	c.L.Lock()
	//for !condition {
	//	c.Wait()
	//}
	fmt.Println("wait before")
	c.Wait()
	fmt.Println("wait end")
	c.L.Unlock()
}

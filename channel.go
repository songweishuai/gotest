package main

import (
	"fmt"
	"time"
)

//func main() {
//	wg := new(sync.WaitGroup)
//	wg.Add(2)
//
//	go func() {
//		defer wg.Done()
//
//		for i := 0; i < 6; i++ {
//			println(i)
//			if i == 3 {
//				runtime.Gosched()
//			}
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		println("Hello World")
//	}()
//
//	wg.Wait()
//}

var ch = make(chan int)

func foo(id int) {
	//fmt.Println("id is:", id)
	ch <- id
}

func main() {
	ch1 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	go func() {
		for val := range ch1 {
			fmt.Println("val is:", val)
			if len(ch1) <= 0 {
				fmt.Println("ch1 is nil")
			}
		}
	}()

	for i := 0; i < 5; i++ {
		go foo(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

	time.Sleep(time.Second * 30)
}

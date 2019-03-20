package main

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

//var ch = make(chan int)
//func foo(id int) {
//	//fmt.Println("id is:", id)
//	ch <- id
//}
//func main() {
//	ch1 := make(chan int, 3)
//
//	ch1 <- 1
//	ch1 <- 2
//	ch1 <- 3
//
//	go func() {
//		for val := range ch1 {
//			fmt.Println("val is:", val)
//			if len(ch1) <= 0 {
//				fmt.Println("ch1 is nil")
//			}
//		}
//	}()
//
//	for i := 0; i < 5; i++ {
//		go foo(i)
//	}
//
//	for i := 0; i < 5; i++ {
//		fmt.Println(<-ch)
//	}
//
//	time.Sleep(time.Second * 30)
//}

// 无缓冲信道
//var done = make(chan bool)
////创建nil chan
////var done chan bool
//var msg string
//
//func aGoroutine() {
//	msg = "hello world"
//	done <- true
//	//close(done)
//}
//
//func main() {
//	go aGoroutine()
//	b := <-done
//	//time.Sleep(time.Second * 3)
//	fmt.Println(msg, b)
//
//	//var testString string = "asd123"
//	//fmt.Println(testString)
//}

// chan 使用技巧 通过channel的缓冲大小控制goroutine数量
var limit chan int = make(chan int, 3)

func main() {
	for _, w := range work {
		// 同时最多只能有3个goroutine执行
		go func() {
			// 超过channel的缓冲大小3以后阻塞
			limit<-1
			w()
			<-limit
		}()
	}

	// 空的管道选择语句，导致阻塞 for{},<-make(chan int)
	select {

	}

}

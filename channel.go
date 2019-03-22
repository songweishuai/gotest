package main

import "fmt"

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
//var limit chan int = make(chan int, 3)
//func main() {
//	for _, w := range work {
//		// 同时最多只能有3个goroutine执行
//		go func() {
//			// 超过channel的缓冲大小3以后阻塞
//			limit<-1
//			w()
//			<-limit
//		}()
//	}
//
//	// 空的管道选择语句，导致阻塞 for{},<-make(chan int)
//	select {
//
//	}
//}

// 带缓冲管道使用技巧
//func main() {
//	done := make(chan int, 10)
//
//	for i := 0; i < 10; i++ {
//		go func() {
//			fmt.Println("你好，世界")
//			done <- 1
//		}()
//	}
//
//	fmt.Println("接受管道信息")
//	for i := 0; i < 10; i++ {
//		<-done
//	}
//}

// 使用sync.WaitGroup
//func main() {
//	var wg sync.WaitGroup
//
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//
//		go func() {
//			fmt.Println("Hello World")
//			wg.Done()
//		}()
//	}
//
//	wg.Wait()
//}

// 生产者，消费者模型
//func Producer(factor int, out chan<- int) {
//	for i := 0; ; i++ {
//		fmt.Println("Producer:", factor)
//		out <- i * factor
//		time.Sleep(time.Second)
//	}
//}
//
//func Consumer(in <-chan int) {
//	for v := range in {
//		fmt.Println(v)
//	}
//}
//
//func main() {
//	ch := make(chan int, 64)
//
//	go Producer(3, ch)
//	go Producer(11, ch)
//	go Consumer(ch)
//
//	//time.Sleep(time.Hour)
//	sig := make(chan os.Signal, 1)
//	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
//	fmt.Println("quit (%v)\n", <-sig)
//}

// select 使用
func main() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:

			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//var total struct {
//	sync.Mutex
//	value int
//}
//
//func Worker(wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for i := 0; i < 100; i++ {
//		total.Lock()
//		total.value += i
//		total.Unlock()
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go Worker(&wg)
//	go Worker(&wg)
//
//	wg.Wait()
//
//	fmt.Println(total.value)
//}

//var total int32
//
//func Worker(wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	var i int32
//	for i = 0; i < 100; i++ {
//		atomic.AddInt32(&total, i)
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	go Worker(&wg)
//	go Worker(&wg)
//
//	wg.Wait()
//
//	fmt.Println(total)
//}

type one struct {
	m    sync.Mutex
	done uint32
}

type www struct {
	a int32
	s string
}

func (o *one) once(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		fmt.Println("此函数已经执行!")
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		f()
		atomic.AddUint32(&o.done, 1)
	}
}

func main() {
	var Once one

	f := func() {
		fmt.Println("hello world")
	}

	Once.once(f)
	Once.once(f)
	Once.once(f)

	fmt.Println("*******************")

	var once sync.Once

	once.Do(f)
	once.Do(f)
	once.Do(f)
	once.Do(f)

	//v := www{10, "asd"}
	//atomic.Value.Store(v)
	//fmt.Println(atomic.Value.Load())

	var config atomic.Value
	v := www{10, "asd"}
	config.Store(v)
	fmt.Println(config.Load())
}

package main

import "fmt"

func minicError(key string) error {
	return fmt.Errorf("minic error:%s",key)
}

func test(){
	fmt.Println("Start test")

	err:=minicError("1")

	defer func() {
		fmt.Println("start defer")

		if err!=nil {
			fmt.Println("defer errors:",err)
		}
	}()
	panic("panic ***********")
	fmt.Println("end test")
}

func main() {
	fmt.Println("main start")
	defer func() {
		fmt.Println("main defer 1")
		fmt.Println(recover())
	}()
	test()
	defer func() {
		fmt.Println("main defer 2")
	}()
	fmt.Println("main end")
}
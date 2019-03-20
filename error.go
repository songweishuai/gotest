package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("defer 1")
			println(err.(string))
		}
	}()

	defer func() {
		fmt.Println("defer 2")
	}()


	defer func() {
		fmt.Println("defer 3")
	}()

	//test()
	panic("panic error")

	log.Fatal()
}

func test() {

	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
			fmt.Println("test function wrapping")
		}
		panic("test panic error")
	}()
}

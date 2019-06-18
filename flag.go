package main

import (
	"flag"
	"fmt"
)

/*
	每一个package是不是都可以拥有参数传入
*/


var a *string = flag.String("a", "", "please input string")
var b *bool = flag.Bool("b", false, "please input bool")

func main() {
	flag.Parse()

	fmt.Println("path is:", *a)
	//ch:=make(chan int)
}

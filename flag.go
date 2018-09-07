package main

import (
	"flag"
	"fmt"
)

var a *string = flag.String("a", "", "please input string")
var b *bool = flag.Bool("b", false, "please input bool")

func main() {
	flag.Parse()

	fmt.Println("path is:", *a)
	//ch:=make(chan int)
}

package main

import "fmt"

func main() {
	type TestInt int

	var a int = 10
	var b int = 10

	if a == b {
		println(true)
	}

	var aa struct {
		x int `a`
	}
	fmt.Printf("%#v\n", aa)
}

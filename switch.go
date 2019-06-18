package main

import "fmt"

func main() {
	x := []int{-1, -2, 3}
	//i := 2

	switch x[1] {
	case x[1]:
		println("a")
		fallthrough
	case 1, 2:
		println("b")
	case 4:
		println("c")
	default:
		println("d")
	}

	println("***********")

	switch {
	case x[0] < 0:
		fmt.Println("aaaa")
	case x[1] > 0:
		println("a")
	case x[1] < 0:
		println("b")
	default:
		println("c")
	}

}

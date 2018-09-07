package main

import "fmt"

func main() {
	//var a = [5]int{1, 2, 3, 4, 5}
	var a [5]int
	fmt.Println(len(a))
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	sum := 0
	for _, num := range a {
		sum += num
	}
	fmt.Println("i: su:", sum)
	fmt.Println("len is", len(a), "haha")

	//var b [3]*int
}

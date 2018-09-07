package main

import "fmt"

func main() {
	//number := []int{1,2,3}
	number := make([]int, 8)

	number = append(number, 0)
	printslice(number)

	number = append(number, 1)
	printslice(number)

	number = append(number, 2, 3, 4, 5, 6, 7, 8, 9)
	printslice(number)

	data := [...]int{1, 2, 3, 4, 10, 10: 0}
	fmt.Println(data)

	s := data[:2:3]
	fmt.Println(s)
	//fmt.Println(cap(s))
	s = append(s, 100,101,102,103,104,105,106,107,108)
	//fmt.Println(cap(s))
	fmt.Println(s)
	fmt.Println(data)
}

func printslice(x []int) {
	for _, a := range x {
		fmt.Println(a)
	}
	fmt.Printf("lend = %d,cap = %d,slice = %v\n", len(x), cap(x), x)
}

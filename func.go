package main

import "fmt"

func main() {
	s1 := test(func() int {
		return 100
	})

	s2 := format(func(s string, x, y int) string  {
		return fmt.Sprintf(s, x, y)
	}, "x=%d,y=%d", 10, 20)

	println(s1, s2)

	println(test1("sum:%d", 1, 2, 3, 4, 5))

	s:=[]int{1,1,1}
	println(test1("sum:%d",s...))
}

func test(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func test1(s string, n ...int) string {
	var x int
	//println(s)
	//println(n)
	for _, i := range n {
		x += i
		//println(i)
	}
	return fmt.Sprintf(s, x)
}

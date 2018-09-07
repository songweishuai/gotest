package main

import "fmt"

func main() {
	fn := func() {
		println("anonymity func")
	}

	fn()

	fns := [](func(x int) int){
		func(x int) int {
			return x + 1
		},
		func(x int) int {
			return x + 2
		},
	}

	println(fns[0](1))
	println(fns[1](2))

	d := struct {
		fn func() string
	}{
		func() string {
			return "Hello,World!"
		},
	}

	println(d.fn())

	/*fc:=make(map[string]string)
	fc["a"]="asd"
	println(fc["a"])*/

	fc := make(chan func() string, 2)
	fc <- func() string {
		return "qwertyu"
	}
	println((<-fc)())

	f := test()
	f()
}

func test() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

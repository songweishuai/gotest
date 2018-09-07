package main

func main() {
	x := []int{-1, -2, 3}
	i := 2

	switch i {
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
	case x[1] > 0:
		println("a")
	case x[1] < 0:
		println("b")
	default:
		println("c")
	}

}

package main

import "fmt"

func getNum() int {
	fmt.Println("get num begin")
	return 0
}

func main() {
	s := `abc`

loop:
	for true {
		for i, n := getNum(), len(s); i < n; i++ {
			println(s[i])
			if i == 0 {
				break loop
			}
		}
	}

	println("********")

	n := len(s)
	for n > 0 {
		n--
		println(s[n])
	}

	println("*********")
	n = len(s)
	for {
		n--
		if n < 0 {
			break
		}
		println(s[n])
	}

	println("*********")
	for i := range s {
		println(s[i])

	}
}

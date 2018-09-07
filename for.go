package main

func main() {
	s := `abc`

	for i, n := 0, len(s); i < n; i++ {
		println(s[i])
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

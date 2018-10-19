package main


func main() {

	panic("panic error")

	//test()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()

	//errors.New()
}

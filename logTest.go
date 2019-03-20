package main

import (
	"errors"
	"fmt"
)

func main() {
	//log.Fatal("#####!!!!@@@####")
	fmt.Println("----------------")
	err := errors.New("1234567890qwer")
	fmt.Println(err)
	//error.Error(err)
}

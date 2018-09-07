package main

import (
	"fmt"
	"test/struct"
	"unsafe"
)

/*
type x struct {
}

func (a *x) test() {
	println(a.test)
}

func main() {
	p := &x{}
	p.test()
}
*/

/*
func main() {
	type User struct {
		UserId   int    `json:"user_id" bson:"user_id"`
		UserName string `json:"user_name" bson:"user_name"`
	}

	u := &User{1, "shuai"}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
}*/


func main() {
	var v = new(P.V)
	var i = (*int32)(unsafe.Pointer(v))
	*i = int32(98)

	var j = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + unsafe.Offsetof(v.J)))
	*j = 99

	fmt.Println("offset i", unsafe.Offsetof(v.I))
	fmt.Println("offset j", unsafe.Offsetof(v.J))

	fmt.Printf("i = %v\n", *i)
	fmt.Printf("j = %v\n", *j)

	fmt.Println(v)

	type a struct {
		x int `a`
	}

	type b struct {
		x int `a`
	}

	var aa a
	var bb b

	aa=bb

}

package main

import (
	"fmt"
)

type person interface {
	say(v interface{})
}

type student struct {
	name string
	age  int
}

func (s student) say(v interface{}) {
	//s1 := v.(string) //reflect
	s1:=string(v)
	fmt.Println(s1)
}

func main() {
	var a int = 98
	fmt.Printf("a = %#v\n", a)
	fmt.Printf("a = %v\n", a)
	fmt.Printf("a = %+v\n", a)
	fmt.Printf("a = %c\n", a)
	fmt.Printf("a = %b\n", a)
	fmt.Printf("a x = 0x%x\n", a)
	fmt.Printf("a X = 0x%x\n", a)
	fmt.Printf("a U = %U\n", a)

	s := student{name: "shuai", age: 20}
	fmt.Printf("s #v = %#v\n", s)
	//fmt.Printf("s v  = %v\n", s)
	//fmt.Printf("s +v = %+v\n", s)
	//fmt.Printf("s T  = %T\n", s)
	////fmt.Printf("s    = %%\n", s)

	//var b = false
	//fmt.Printf("b t  = %t\n", b)

	var str = "song wei shuai"
	fmt.Printf("str = %s\n", str)
	fmt.Printf("str = %q\n", str)

	ptr := new([]int)
	fmt.Printf("ptr = %p\n", ptr)
	fmt.Printf("ptr = %v\n", ptr)

	ch := make(chan int)
	fmt.Printf("ch = %v\n", ch)

	s1 := fmt.Sprintf("%[2]d %[1]d", 111, 222)
	fmt.Println(s1)

	s2 := str + s1
	fmt.Println(s2)

	s3 := fmt.Sprintf("%6.1f", 12.0)
	fmt.Println(s3)

	s.say("hello world")
}

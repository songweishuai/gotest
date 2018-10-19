package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
type Stinger interface {
	String() string
}

type Printer interface {
	Stinger
	Print()
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d,%s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func main() {
	//var t Printer = &User{1, "Tom"}
	//var t *User=&User{1,"Tom"}
	//u := User{1, "tom"}
	var t Printer=&User{1,"Tom"}
	//var t Printer = u
	t.Print()
}*/

/*
type Tester struct {
	s interface{
		String() string
	}
}

type User struct {
	id int
	name string
}

func (self *User)String() string{
	return fmt.Sprintf("user %d,%s",self.id,self.name)
}

func main() {
	t:=Tester{&User{1,"Tom"}}
	fmt.Println(t.s.String())
}
*/

/*
type User struct {
	id   int
	name string
}

func (self User) String() string {
	return fmt.Sprintf("%d %s", self.id, self.name)
}

func main() {
	u := User{1, "Tom"}
	var vi, pi interface{} = u, &u

	//vi.(User).name = "jack"
	pi.(*User).name = "Jack"

	fmt.Printf("%v\n", vi.(User))
	fmt.Printf("%v\n", pi.(*User))

	var o interface{} = l{1, "tom"}
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
		fmt.Println(ok)
	}

	u := o.(*User)//o=&User{}
	fmt.Println(u)
}*/

/*
type User struct {
	id   int
	name string
}

func main() {
	var o interface{} = &User{1, "shuai"}

	switch v := o.(type) {
	case nil:
		fmt.Println("nil")
	case fmt.Stringer:
		fmt.Println(v)
	case func() string:
		fmt.Println(v())
	case *User:
		fmt.Printf("%d %s\n", v.id, v.name)
	default:
		fmt.Println("unknow")
	}
}*/

//type Stringer interface {
//	String() string
//}
//
//type Printer interface {
//	String() string
//	//Stringer
//	Print()
//}
//
//type User struct {
//	id int
//	name string
//}
//
//func (self *User)String() string{
//	return fmt.Sprintf("%d,%v",self.id,self.name)
//}
//
//func (self *User)Print()  {
//	fmt.Println(self.String())
//}
//
//func main() {
//	var o Printer=&User{1,"tom"}
//	var s Stringer=o
//	fmt.Println(s.String())
//}



type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	//return p.Write(bytes.ToUpper(data))
	return p.Writer.Write(bytes.ToUpper(data))
}

type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello world")

	fmt.Fprintln(os.Stdout, UpperString("hello world"))
}

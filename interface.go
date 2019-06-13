package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// 通过实现io.Writer接口,将字母转化为大写
type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

type UpString string

func (p UpString) String() string {
	return strings.ToUpper(string(p))
}

func main() {
	out := UpperWriter{os.Stdout}
	fmt.Fprintln(&out, "hello,world")
	fmt.Println(UpString("songweishuai"))
	fmt.Fprintln(os.Stdout, UpString("songweishuai"))
	//runtime.Error()
	//generator.Generator{}
	//proto.Message()
}

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

/*匿名接口用作变量类型或结构体成员*/
//type Tester struct {
//	s interface{
//		String() string
//	}
//}
//
//type User struct {
//	id int
//	name string
//}
//
//func (self *User)String() string{
//	return fmt.Sprintf("user %d,%s",self.id,self.name)
//}
//
//func main() {
//	t:=Tester{&User{1,"Tom"}}
//	fmt.Println(t.s.String())
//}

/*修改接口对象的数据需要使用指针类型*/
//type User struct {
//	id   int
//	name string
//}
//
//func (self User) String() string {
//	return fmt.Sprintf("%d %s", self.id, self.name)
//}
//
//func main() {
//	u := User{1, "Tom"}
//	var vi, pi interface{} = u, &u
//
//	//vi.(User).name = "jack"
//	pi.(*User).name = "Jack"
//
//	fmt.Printf("%v\n", vi.(User))
//	fmt.Printf("%v\n", pi.(*User))
//}

/*接口类型判断 assertion*/
//type User struct {
//	id   int
//	name string
//}
//
//func main() {
//	var o interface{} = User{1, "shuai"}
//
//	switch v := o.(type) {
//	case nil:
//		fmt.Println("nil")
//	case fmt.Stringer:
//		fmt.Println(v)
//	case func() string:
//		fmt.Println(v())
//	case User:
//		fmt.Printf("%d %s\n", v.id, v.name)
//	default:
//		fmt.Println("unknow")
//	}
//}

/*接口转型，大类型可以转换为小类型。*/
/*方法接受者为指针，则必须使用指针；如果方法的接受者为变量，使用指针赋值也可以。*/
//type Stringer interface {
//	String() string
//}
//
//type Printer interface {
//	//String() string
//	Stringer
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

/*继承了接口规范*/
//type UpperWriter struct {
//	io.Writer
//}
//
//func (p *UpperWriter) Write(data []byte) (n int, err error) {
//	//return p.Write(bytes.ToUpper(data))
//	return p.Writer.Write(bytes.ToUpper(data))
//}
//
//type UpperString string
//
//func (s UpperString) String() string {
//	return strings.ToUpper(string(s))
//}
//
//func main() {
//	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello world")
//
//	fmt.Fprintln(os.Stdout, UpperString("hello world"))
//}

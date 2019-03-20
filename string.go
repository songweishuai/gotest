package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "Hello World"
	s = s + "songweishuai"
	fmt.Println(s, len(s))

	//s = fmt.Sprintf("")
	s4 := "songweishuai"

	s1 := s[:5]
	s2 := s[7:]
	s3 := s[1:5]
	s5 := "songweishuai"

	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s4)).Data)
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s5)).Data)

	fmt.Println("s1 is ", s1)
	fmt.Println("s2 is ", s2)
	fmt.Println("s3 is ", s3)

	/*Unicode Code Point*/
	fmt.Printf("%T\n", 'a')

	// 字符串的只读属性禁止程序修改字符串底层对应的字节数组元素。但是可以把字符串转换为rune和byte类型进行修改。
	bs := []rune(s)
	bs[0] = 'w'
	fmt.Printf("s:%s\n", string(bs))

	bss := []byte(s)
	bss[1] = 'w'
	fmt.Printf("bss:%s\n", string(bss))
	s = string(bss)
	fmt.Printf("s:%s\n", s)

	//utf8.DecodeRuneInString()
	fmt.Println("\xe4\xb8\x96")
	fmt.Println("\xe7\x95\x8c")
	fmt.Printf("%#v\n", []byte("\xe7\x95\x8c"))
	fmt.Println("\xe4\x00\x00\xe7\x95\x8casdfghjkl")

	//errors.New()

	// rune[]和string类型的相互转换.可以发现[]rune其实就是[]int32类型，rune其实就是int32的别名，并不是重新定义的类型。
	fmt.Printf("%#v\n", []rune("abc"))
	fmt.Printf("%#v\n", rune('a'))

	//var 宋 = 10
	//fmt.Println(宋)

	var a int
	fmt.Println(a)
}

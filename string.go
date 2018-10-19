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
	fmt.Println("\xe4\x00\x00\xe7\x95\x8casdfghjkl")

	//errors.New()
}

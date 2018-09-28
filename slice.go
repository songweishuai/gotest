package main

import (
	"fmt"
	"reflect"
	"unsafe"
	)

type Slice struct {
	ptr unsafe.Pointer
	len int
	cap int
}

func main() {
	//array := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := array[:5]

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := append(s1, 10)
	slice1 := (*Slice)(unsafe.Pointer(&s1))
	fmt.Println("ptr:", slice1.ptr, "len", len(s1), "cap:", cap(s1))
	slice2 := (*Slice)(unsafe.Pointer(&s2))
	fmt.Println("ptr:", slice2.ptr, "len", len(s2), "cap:", cap(s2))

	if slice1 == slice2 {
		fmt.Println("slice1 == slice2")
	} else {
		fmt.Println("slice1 != slice2")
	}

	s3 := make([]int, 5, 6)
	for i := 0; i < 5; i++ {
		s3[i] = i
	}

	//s4 := append(s3, 5)
	s4 := s3[:0]
	fmt.Println(s4)

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s3)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s4)).Data)

	//regexp.MustCompile()
	append()
}

//func main() {
//	//number := []int{1,2,3}
//	number := make([]int, 8)
//
//	number = append(number, 0)
//	printslice(number)
//
//	number = append(number, 1)
//	printslice(number)
//
//	number = append(number, 2, 3, 4, 5, 6, 7, 8, 9)
//	printslice(number)
//
//	data := [...]int{1, 2, 3, 4, 10, 10: 0}
//	fmt.Println(data)
//
//	s := data[:2:3]
//	fmt.Println(s)
//	//fmt.Println(cap(s))
//	s = append(s, 100,101,102,103,104,105,106,107,108)
//	//fmt.Println(cap(s))
//	fmt.Println(s)
//	fmt.Println(data)
//}
//
//func printslice(x []int) {
//	for _, a := range x {
//		fmt.Println(a)
//	}
//	fmt.Printf("lend = %d,cap = %d,slice = %v\n", len(x), cap(x), x)
//}

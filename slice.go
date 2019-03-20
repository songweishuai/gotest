package main

import (
	"fmt"
	"unsafe"
)

type Slice struct {
	ptr unsafe.Pointer
	len int
	cap int
}

func main() {
	// 切片基础信息
	////nil切片定义
	//var a []int
	////非nil切片定义,但是len和cap都为0
	//var a = []int{}
	//fmt.Println("len(a):", len(a))
	//fmt.Println("cap(a):", cap(a))
	//if a == nil {
	//	fmt.Println("a == nil")
	//}

	// 切片尾部添加数据
	//var a []int
	//a = append(a, 1)
	//a = append(a, 1, 2, 3)
	//a = append(a, []int{1, 2, 3}...)
	//fmt.Println(a)

	// 切片头部添加数据，在开头添加数据一般都会导致内存的从新分配，性能要比在尾部添加差。
	//var a = []int{1, 2, 3}
	//a = append([]int{0}, a...)
	//a = append([]int{-3, -2, -1}, a...)
	//fmt.Println(a)

	// 在切片中间添加数据，利用append函数返回也是切片使用链式操作
	//var a = []int{0, 1, 2, 3}
	//a = append(a[:1], append([]int{-1, -2, -3}, a[1:]...)...)
	//fmt.Println(a)

	// 删除切片开头元素
	//a := []int{1, 2, 3}
	//b := a[1:]
	//fmt.Println("b:", b)
	//c := copy(a, a[1:])
	//fmt.Println("c:", c)
	//a = a[:copy(a, a[1:])]
	//fmt.Println("a:", a)

	// 删除中间元素
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a:", a)
	//a = a[:3+copy(a[3:], a[3+1:])]
	b := copy(a[3:], a[3+1:])
	fmt.Println("b:", b)
	a = a[:3+b]
	fmt.Println("a:", a)

	//x := []int{2, 3, 5, 7, 11}
	//y := x[1:2]
	//fmt.Println("x:", x)
	//fmt.Println("y:", y, len(y), cap(y))
}

//func main() {
//	//array := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	//s1 := array[:5]
//
//	//var a []int
//	//a = append(a, 1)
//	////a = append(a, 1, 2, 3)
//	//a = append(a, []int{1, 2, 3}...)
//	//fmt.Println(a)
//	//os.Exit(1)
//
//	// 测试当切片容量不足的时候会从新分配内存
//	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	s2 := append(s1, 10)
//	slice1 := (*Slice)(unsafe.Pointer(&s1))
//	// 打印结果切片的数据指针部一样
//	fmt.Println("ptr:", slice1.ptr, "len", len(s1), "cap:", cap(s1))
//	slice2 := (*Slice)(unsafe.Pointer(&s2))
//	fmt.Println("ptr:", slice2.ptr, "len", len(s2), "cap:", cap(s2))
//	if slice1 == slice2 {
//		fmt.Println("slice1 == slice2")
//	} else {
//		fmt.Println("slice1 != slice2")
//	}
//
//
//
//	s3 := make([]int, 5, 6)
//	for i := 0; i < 5; i++ {
//		s3[i] = i
//	}
//
//	//s4 := append(s3, 5)
//	s4 := s3[:0]
//	fmt.Println(s4)
//
//	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s3)).Data)
//	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s4)).Data)
//
//	//regexp.MustCompile()
//	//append()
//}

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

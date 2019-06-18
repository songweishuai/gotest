package main

import (
	"fmt"
	"reflect"
)

//func main() {
//	type TestInt int
//
//	var a int = 10
//	var b int = 10
//	var f float64 = 10.01
//	c := int32(f)
//	fmt.Println("c:", c)
//
//	if a == b {
//		println(true)
//	}
//
//	var aa struct {
//		x int `json:"width"`
//	}
//	fmt.Printf("%#v\n", aa)
//
//	fl := []float64{1.0, 2.1, 2.3, 2, 5, 3.6, 4.9}
//	fmt.Println(fl)
//	SortFloat64FastV2(fl)
//	fmt.Println(fl)
//}
//
//func SortFloat64FastV2(a []float64) {
//	var c []int
//
//	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
//	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
//
//	*cHdr = *aHdr
//
//	sort.Ints(c)
//}

//const max  = 100
//
//
//// go语言是强类型语言,不同类型的变量不可以互相赋值
//type myint int
//
//func main() {
//	var a myint = 0
//	var b int = 1
//	// 必须使用强制类型转换
//	a = myint(b)
//	fmt.Println(a, b)
//}

// 类型判断
func main() {
	var a interface{}
	//a = "asd"
	//a=10
	a = map[interface{}]interface{}{
		"a": "b",
	}

	switch t := a.(type) {
	case int:
		fmt.Println(t)
		fmt.Println("my type is int")
	case string:
		fmt.Println("my type is string")
	case map[interface{}]interface{}:
		fmt.Println("my type is map")
	}
	fmt.Println(fmt.Sprintf("%T", a))
	fmt.Printf("%T\n", a)

	myself(1, 2, "3", 4, 5, 6)
}

func myself(k int, val ...interface{}) {
	fmt.Println("k:", k)
	fmt.Println(val...)
	fmt.Println("不解包:", val)
	fmt.Println(reflect.TypeOf(val[0]))

	switch reflect.TypeOf(val[0]).String() {
	case "int":
		fmt.Println("My type is int")
	case "string":
		fmt.Println("My type is string")
	}

	switch reflect.TypeOf(val[1]).String() {
	case "int":
		fmt.Println("My type is int")
	case "string":
		fmt.Println("My type is string")
	}

	fmt.Println(reflect.ValueOf(val))
}
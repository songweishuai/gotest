package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func main() {
	type TestInt int

	var a int = 10
	var b int = 10
	var f float64 = 10.01
	c := int32(f)
	fmt.Println("c:", c)

	if a == b {
		println(true)
	}

	var aa struct {
		x int `a`
	}
	fmt.Printf("%#v\n", aa)

	fl := []float64{1.0, 2.1, 2.3, 2, 5, 3.6, 4.9}
	fmt.Println(fl)
	SortFloat64FastV2(fl)
	fmt.Println(fl)
}

func SortFloat64FastV2(a []float64) {
	var c []int

	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))

	*cHdr = *aHdr

	sort.Ints(c)
}

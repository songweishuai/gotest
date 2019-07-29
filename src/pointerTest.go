package main

import (
	"fmt"
	P "struct"
	"unsafe"
)

func main() {
	var p = new(P.V)
	var i = (*int32)(unsafe.Pointer(p))

	*i = int32(88)

	fmt.Println(unsafe.Sizeof(0))
	var j = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(unsafe.Sizeof(0))))
	*j = 66

	p.PutI()
	p.PutJ()
}

package P

import (
	"fmt"
	"unsafe"
)

type V struct{
	b byte
	I int32
	J int64
}

func (this V)PutI()  {
	fmt.Println("this.i=%v",this.I)
}

func (this V)PutJ(){
	fmt.Println("this.j=%v",this.J)
}

func init()  {
	var v =new(V)
	fmt.Println("szie=%d",unsafe.Sizeof(*v))
}
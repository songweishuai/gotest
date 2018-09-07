package main

import (
	"time"
	"fmt"
)

func main() {
	t0:=time.Now()
	fmt.Println("t0 is:",t0)

	t1:=time.Now()
	fmt.Println("t1 is:",t1)

	fmt.Println("the call took %v to run\n",t1.Sub(t0))
}

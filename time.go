package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	fmt.Println("t0 is:", t0)

	//time.Sleep(3 * time.Second)
	fmt.Println(time.Since(t0))

	t1 := time.Now()
	fmt.Println("t1 is:", t1)

	//fmt.Println("befor:",t1.Before())
	t2 := t1.Add(1000000000000000)
	fmt.Println("t2 is:", t2)
	//fmt.Println("Until", time.Until(t1))
	fmt.Println("the call took %v to run\n", t1.Sub(t0))


	// 计算时区
	fmt.Println("计算时区")
	staticStr := "2018-03-09 10:00:00"
	var timeFomatLoc *time.Location
	timeFomatLoc,err:=time.LoadLocation("Local")
	fmt.Println(timeFomatLoc)
	Format := "2006-01-02 15:04:05"
	t,err:=time.ParseInLocation(Format,staticStr,timeFomatLoc)
	if err == nil {
		fmt.Println(t)
	}
}

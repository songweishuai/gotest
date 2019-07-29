package main

import (
	"fmt"
	"time"
)

func main() {
	var fillInterval = time.Second*3
	var capacity = 100
	var tokenBucket = make(chan struct{}, capacity)

	fillToken := func() {
		ticker := time.NewTicker(fillInterval)
		for {
			select {
			case val := <-ticker.C:
				fmt.Println("val:", val)
				select {
				case tokenBucket <- struct{}{}:
				default:

				}
				//fmt.Println("current token cnt:", len(tokenBucket), time.Now())
			}
		}
	}

	go fillToken()
	time.Sleep(time.Hour)
	//rand.Perm()
	//parser.Trace
}

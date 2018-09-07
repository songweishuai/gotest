package main

import (
	"fmt"
	"os"
	"test/checkErr"
)

func main() {
	file, err := os.Open("1.txt")
	checkErr.CheckErr(err, true)
	defer file.Close()

	fileInfo, err := file.Stat()
	checkErr.CheckErr(err, true)
	//fmt.Println("fileInfo is:", fileInfo)
	fmt.Printf("%#v\n", fileInfo.Size())

	var str string
	var fileLen int64
	for {
		data := make([]byte, 10)
		count, err := file.Read(data)
		checkErr.CheckErr(err, true)
		//fmt.Println("count:", count)
		//fmt.Printf("data:%s\n", data)
		str = str + string(data)
		fileLen += int64(count)
		if fileLen >= fileInfo.Size() {
			fmt.Println("read over")
			break
		}
	}

	fmt.Println("file info:", str)

	fmt.Println(os.Hostname())
}

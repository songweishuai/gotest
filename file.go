package main

import (
	"fmt"
	"os"
)

//func main() {
//	file, err := os.Open("1.txt")
//	checkErr.CheckErr()
//	defer file.Close()
//
//	fileInfo, err := file.Stat()
//	checkErr.CheckErr(err, true)
//	//fmt.Println("fileInfo is:", fileInfo)
//	fmt.Printf("%#v\n", fileInfo.Size())
//
//	var str string
//	var fileLen int64
//	for {
//		data := make([]byte, 10)
//		count, err := file.Read(data)
//		checkErr.CheckErr(err, true)
//		//fmt.Println("count:", count)
//		//fmt.Printf("data:%s\n", data)
//		str = str + string(data)
//		fileLen += int64(count)
//		if fileLen >= fileInfo.Size() {
//			fmt.Println("read over")
//			break
//		}
//	}
//
//	fmt.Println("file info:", str)
//
//	fmt.Println(os.Hostname())
//}

func main() {
	file, err := os.Create("./1.txt")
	if err != nil {
		fmt.Println(err)
	}
	data := make([]byte, 128)
	n, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("n:", n)
}

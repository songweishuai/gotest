package main

import (
	"fmt"
	"github.com/chai2010/errors"
	"io/ioutil"
	"log"
)

//func main() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println("defer 1")
//			println(err.(string))
//		}
//	}()
//
//	defer func() {
//		fmt.Println("defer 2")
//	}()
//
//
//	defer func() {
//		fmt.Println("defer 3")
//	}()
//
//	//test()
//	panic("panic error")
//
//	log.Fatal()
//}
//
//func test() {
//
//	defer func() {
//		if err := recover(); err != nil {
//			println(err.(string))
//			fmt.Println("test function wrapping")
//		}
//		panic("test panic error")
//	}()
//}

func loadConfig() error {
	_, err := ioutil.ReadFile("/path/to/file")
	if err != nil {
		return errors.Wrap(err, "read failed")
	}
	return nil
}

func setup() error {
	err := loadConfig()
	if err != nil {
		return errors.Wrap(err, "invalid config")
	}
	return nil
}

func main() {
	if err := setup(); err != nil {
		// 获取包装流程
		for i, e := range err.(errors.Error).Wraped() {
			fmt.Printf("wraped(%d): %v\n", i, e)
		}

		// 函数调用堆栈信息
		for i, x := range err.(errors.Error).Caller() {
			fmt.Printf("caller:%d: %s\n", i, x.FuncName)
		}

		// 将错误信息编码为json字符串
		errors.ToJson(err)

		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"reflect"
)

/*
	接口式函数编程
	1.只适用于只有一个方法的接口.
	2.将函数定义为类型,参数和返回值跟接口方法一样.
	3.类型实现接口方法
	4.使用类型的强制类型转换,将函数转换为该类型
	接口式函数意义:
	1.可以定义不同于接口方法名字的函数,使函数名定义更有意义.
	2.不用定义新类型,就可以实现接口.
*/

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, f Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			f.Do(k, v)
		}
	}
}

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

func selfInfo(k, v interface{}) {
	fmt.Println("大家好我叫:%s,今年%s岁", k, v)
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

func main() {
	var m map[interface{}]interface{} = map[interface{}]interface{}{
		"张三": "10",
		"李四": "11",
		"王五": "12",
	}

	Each(m, HandlerFunc(func(k, v interface{}) {
		fmt.Println("大家好我叫%s,今年%s", k, v)
	}))

	//Each(m, HandlerFunc(selfInfo))

	myself(1, 2, "3", 4, 5, 6)


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
}

package main

import (
	"fmt"
)

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

type welcome string

func (w welcome) selfinfo(k, v interface{}) {
	fmt.Printf("%s,my name is %s,%d\n", w, k, v)
}


func main() {

	//person:=make(map[interface{}]interface{})
	//person["zs"]=20
	//person["ls"]=28
	//person["ww"]=26

	person := map[interface{}]interface{}{
		"zs": 20, "ls": 28, "ww": 26,
	}

	var w welcome = "welcome"

	//Each(person, HandlerFunc(w.selfinfo))
	//Each(person,HandlerFunc(func(k, v interface{}) {
	//	fmt.Printf("%s,my name is %s,%d\n", w, k, v)
	//
	//}))
	//EachFunc(person, w.selfinfo)
	//EachFunc(person, func(k, v interface{}) {
	//	fmt.Printf("%s,my name is %s,%d\n", w, k, v)
	//})
}

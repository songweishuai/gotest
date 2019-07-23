package main

import (
	"fmt"
)

//func main() {
//	m := map[string]int{"a": 1, "b": 2}
//
//	for key, value := range m {
//		println(key, value)
//	}
//}

type Point struct {
	x int
	y int
}

func main() {
	r := make([]*Point, 4)
	d := []Point{
		{1, 3},
		{3, 3},
		{3, 48},
		{8, 2},
	}
	for _, v := range d {
		fmt.Println(v)
		r = append(r, &v)
	}

	for _, v := range r {
		fmt.Println(v)
	}
}

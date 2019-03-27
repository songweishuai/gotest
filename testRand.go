package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())

	// 返回0-10的随机序列
	fmt.Println(rand.Perm(10))
	fmt.Println(rand.Perm(10))
	fmt.Println(rand.Perm(10))
	fmt.Println(rand.Perm(10))
	fmt.Println(rand.Perm(10))
	fmt.Println(rand.Perm(10))
}

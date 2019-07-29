package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func main() {
//	// 初始化随机种子
//	rand.Seed(time.Now().UnixNano())
//
//	// 返回0-10的随机序列
//	fmt.Println(rand.Perm(10))
//	fmt.Println(rand.Perm(10))
//	fmt.Println(rand.Perm(10))
//	fmt.Println(rand.Perm(10))
//	fmt.Println(rand.Perm(10))
//	fmt.Println(rand.Perm(10))
//
//	rand.Int()
//}

// 初始化随机种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

// 创建索引切片，随机索引切片,使用两两交换方式
func shuffle1(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}

// fisher-yates算法 创建索引切片，随机索引切片，使用随机选取一个放到最后
func shuffle2(slice []int) {
	for i := len(slice) - 1; i > 0; i-- {
		a := rand.Intn(i)
		slice[i], slice[a] = slice[a], slice[i]
	}
}

func main() {
	var cnt1 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		var s1 = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle1(s1)
		cnt1[s1[0]]++
	}

	var cnt2 = map[int]int{}
	for i := 0; i < 1000000; i++ {
		var s2 = []int{0, 1, 2, 3, 4, 5, 6}
		shuffle2(s2)
		cnt2[s2[0]]++
	}

	fmt.Println(cnt1, "\n", cnt2)
}

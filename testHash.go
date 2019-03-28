package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

var bucketSize uint64 = 10

func main() {
	var bucketMap = map[uint64]int{}

	for i := 15000000000; i < 15000000000+1000000; i++ {
		//hashInt := murmur3.Sum64([]byte(fmt.Sprint(i))) % bucketSize
		// 平均散列到10个桶中
		hashInt := murmur32(fmt.Sprint(i)) % bucketSize
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
}

func murmur32(p string) uint64 {
	return murmur3.Sum64([]byte(p))
}

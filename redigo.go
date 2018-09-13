package main

import (
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
	"reflect"
	"time"
)

var del = flag.Bool("del", false, "delete key?")

func main() {
	flag.Parse()

	con, err := redis.DialTimeout("tcp", "127.0.0.1:6379", 3*time.Second, 3*time.Second, 3*time.Second)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer con.Close()

	exist, err := con.Do("exists", "k1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if exist == int64(1) && *del == true {
		del, err := con.Do("del", "k1")
		if err != nil {
			fmt.Println("del err is:", err)
			os.Exit(1)
		}
		fmt.Println(reflect.TypeOf(del), del)

		os.Exit(0)
	}
	fmt.Println("exist is:", exist, reflect.TypeOf(exist))

	getset, err := con.Do("getset", "k2", "123")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("getset", getset, reflect.TypeOf(getset))
	fmt.Println(redis.String(getset, err))

	len, err := con.Do("strlen", "k2")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("len=", len, reflect.TypeOf(len))

	re, err := con.Do("set", "k1", 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(re)

	/*set expire time second*/
	expire, err := con.Do("expire", "k1", 10)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("expire is:", expire)

	re1, err := con.Do("get", "k1")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(re1)
	fmt.Println(reflect.TypeOf(re1))
	fmt.Println(redis.Int(re1, err))
}

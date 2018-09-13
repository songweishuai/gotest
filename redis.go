package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func ExampleNewClient() *redis.Client {
	//return nil
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	fmt.Println(pong, err)
	return client
}

func ExampleClient(client *redis.Client) bool {
	defer func() {
		err := recover()
		fmt.Println("recover panic is:", err)
	}()

	if client == nil {
		panic("client is null")
	}

	v, e := client.Exists("key").Result()
	if e != nil {
		panic(e)
	}

	if v > 0 {
		fmt.Println(v)
		return false
	}

	if v <= 0 {
		fmt.Println(v)
		return false
	}

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("val is:", val)

	return true
}

func main() {
	client := ExampleNewClient()
	if client != nil {
		defer client.Close()
	}

	ExampleClient(client)

	redis
}

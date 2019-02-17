package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	//写入redis
	err := client.Set("test", []byte("hello beifeng"))
	if err != nil {
		panic(err)
	}
	//读取redis
	res, err := client.Get("test")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(res))
}

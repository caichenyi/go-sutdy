package main

import (
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	//写入redis
	f := make(map[string]interface{})
	f["name"] = "zhangsan"
	f["age"] = 12
	f["sex"] = "male"

	_, err := client.Zadd("test_zset", []byte("beifeng"), 100)
	if err != nil {
		panic(err)
	}
}

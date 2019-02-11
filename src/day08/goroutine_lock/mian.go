package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m = make(map[int]uint64)
	lock sync.Mutex //互斥所
)

type task struct {
	n int
}

func calc(t *task) {
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++{
		sum *= uint64(i)
	}
	lock.Lock()   //加锁
	m[t.n] = sum
	lock.Unlock()  //解锁
}

func main() {
	for i := 1; i < 50; i++ {
		t := &task{n: i}
		go calc(t)
	}
	lock.Lock()   //加锁
	time.Sleep(5 * time.Second)
	for k, v := range m {
		fmt.Printf("%d! = %v\n", k, v)
	}
	lock.Unlock()  //解锁
}

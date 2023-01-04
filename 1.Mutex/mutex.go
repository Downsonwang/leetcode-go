package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var count int
var Mutex sync.Mutex
var mutex sync.RWMutex

func Printer(s string) {
	//锁定以下的共享资源
	Mutex.Lock()

	for _, data := range s {
		fmt.Printf("%c ", data)
	}
	fmt.Println("")

	defer Mutex.Unlock()
}
func write(n int) {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("写 goroutine %d 正在写数据 \n", n)
	mutex.Lock()
	num := rand.Intn(500)
	count = num
	fmt.Printf("写 goroutine %d 写数据结束 %d\n", n, num)
	mutex.Unlock()
}
func read(n int) {
	mutex.RLock()
	fmt.Printf("读 goroutine %d 正在读取数据 \n", n)
	num := count
	fmt.Printf("读 goroutine %d 读取数据结束 %d\n", n, num)
	mutex.RUnlock()
}

func main() {
	//go Printer("Hello")
	//go Printer("World")
	for i := 0; i < 10; i++ {
		go read(i + 1)
	}
	for i := 0; i < 10; i++ {
		go write(i + 1)
	}

	time.Sleep(time.Second * 5)
}

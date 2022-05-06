package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var couter int = 0
var sum int = 0

func Add(a int, b int, lock *sync.Mutex) {
	fmt.Println("InAdd")
	sum = sum + a + b
	lock.Lock()
	couter++
	lock.Unlock()
}
func Sub(a int, b int, lock *sync.Mutex) {
	fmt.Println("InSub")
	sum = sum + a - b

	lock.Lock()
	couter++
	lock.Unlock()
}
func main() {
	start := time.Now()
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Add(i, i+1, lock)
	}
	for {
		lock.Lock()
		c := couter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}

	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)

}

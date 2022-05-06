package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var couter int = 0

func Say(s string, lock *sync.Mutex) {
	for i := 0; i < 3; i++ {

		fmt.Println(s)
		lock.Lock()
		couter++
		lock.Unlock()

	}

}
func Do(s string) {
	fmt.Println(s)

}
func main() {
	start := time.Now()
	lock := &sync.Mutex{}
	go Say("step 1", lock)
	/* for {
		if couter >= 3 {
			break
		}
	} */
	runtime.Gosched()
	Do("step 2")
	runtime.Gosched()
	Do("step 3")
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)

}

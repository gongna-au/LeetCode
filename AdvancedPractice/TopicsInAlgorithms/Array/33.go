package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.NewTimer(6 * time.Second)
	fmt.Println("waiting")
	go func() {

		fmt.Println("time out")
	}()
	stop := time1.Stop()
	if stop {
		fmt.Println("stop")
	}

}

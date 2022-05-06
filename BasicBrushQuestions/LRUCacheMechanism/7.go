package main

import (
	"fmt"
	"time"
)

func Work(stopChan chan struct{}) {

	t := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-stopChan:
			fmt.Println("main ask to stop")
			return
		case <-t.C:
			fmt.Println("Working")
		}

	}

	return

}
func main() {
	stopChan := make(chan struct{})
	Work(stopChan)
	time.Sleep(time.Second * 3)
	close(stopChan)
	fmt.Println("end")

}

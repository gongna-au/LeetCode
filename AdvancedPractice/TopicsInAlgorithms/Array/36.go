package main

import (
	"fmt"
	"sync"
)

func Work(s string) {
	for i := 0; i < 4; i++ {
		fmt.Println("i=", i, s)
	}

}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Work("test")

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		Work("done")

	}()
	wg.Wait()

	fmt.Println("end")

}

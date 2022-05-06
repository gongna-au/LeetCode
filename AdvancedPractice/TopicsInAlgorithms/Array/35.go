package main

import (
	"fmt"
	//"sync"
)

func Work(s string, ch chan bool) {
	for i := 0; i < 4; i++ {
		fmt.Println("i=", i, s)
	}
	ch <- true
}
func main() {
	//var wg sync.WaitGroup
	ch := make(chan bool)
	ch1 := make(chan bool)

	go func() {

		Work("test", ch)

	}()

	go func() {
		Work("done", ch1)

	}()

	k, _ := <-ch
	if k == true {
		fmt.Println("work1 has been finished you can continue next work")

	}
	k, _ = <-ch1
	if k == true {
		fmt.Println("work2 has been finished you can continue next work")

	}

	fmt.Println("end")

}

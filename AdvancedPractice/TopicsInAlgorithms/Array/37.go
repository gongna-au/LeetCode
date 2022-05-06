package main

import (
	"fmt"
	"sync"
	// "time"
)

type hander func(string)

func Sing(s string) {
	for i := 0; i < 4; i++ {
		fmt.Println("sing", s)
	}

}
func Say(s string) {
	for i := 0; i < 4; i++ {
		fmt.Println("saying", s)
	}

}

func main() {

	tasks := []hander{
		func(s string) {
			Sing(s)
		},
		func(s string) {
			Say(s)

		},
	}
	var wg sync.WaitGroup
	wg.Add(2)
	for i, v := range tasks {
		fmt.Println(i)
		// go func() {
		s := "bird"
		v(s)
		wg.Done()
		// }()

	}

	//fmt.Println(1)

	// time.Sleep(8 * time.Second)
	wg.Wait()
	fmt.Print("end")

}

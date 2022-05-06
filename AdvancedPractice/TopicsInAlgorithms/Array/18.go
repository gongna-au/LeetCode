package main

import "fmt"

func main() {
	ch := make(chan int)
	defer close(ch)

	go func() {
		for {
			k, _ := <-ch
			fmt.Println(k)
		}

	}()
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

}

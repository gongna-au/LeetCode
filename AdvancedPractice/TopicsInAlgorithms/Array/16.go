package main

import "fmt"

func Producer(out chan<- int) {

	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)

}
func Consumer(in <-chan int) {
	for {
		if k, ok := <-in; ok {
			fmt.Println("k", k)
		} else {
			break
		}

	}

}
func main() {
	ch := make(chan int, 5)
	go Producer(ch)
	Consumer(ch)

}

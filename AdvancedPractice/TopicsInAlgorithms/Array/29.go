package main

import "fmt"

func Producer(slice []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range slice {
			ch <- v
		}

	}()

	return ch
}
func Consumer(ch <-chan int) <-chan int {
	ch1 := make(chan int)

	go func() {
		defer close(ch1)
		for k := range ch {
			ch1 <- k * k
		}

	}()
	return ch1

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	ch := Producer(array)
	endCh := Consumer(ch)
	for k := range endCh {
		fmt.Println(k)
	}

}

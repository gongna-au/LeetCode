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
	ch := make(chan int)
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	go func() {
		defer close(ch)
		for _, v := range slice {
			ch <- v
		}

	}()
	for k := range ch {
		fmt.Println(k)
	}
}

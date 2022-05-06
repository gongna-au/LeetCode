package main

import (
	"fmt"
	"sync"
	//"time"
)

type temp struct {
	name string
	age  int
}

func Step1(s ...int) <-chan int {
	result := make(chan int)
	go func() {
		for _, v := range s {
			result <- v

		}
		close(result)

	}()

	return result
}
func Step2(in <-chan int) <-chan int {
	result := make(chan int)
	go func() {
		for v := range in {
			result <- v * v
		}
		close(result)

	}()

	return result
}
func Step3(ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	result := make(chan int)

	output := func(c <-chan int) {
		for v := range c {
			result <- v
		}
		wg.Done()

	}
	wg.Add(len(ch))
	for _, v := range ch {

		go output(v)

	}

	wg.Wait()
	close(result)

	return result
}

func (t *temp) Test() {
	t.name = "1234"
	fmt.Println(t.name)
	fmt.Println("test success !")

}

func main() {
	ch := Step1(1, 2, 3, 4)
	temp1 := Step2(ch)
	temp2 := Step2(ch)
	for n := range Step3(temp1, temp2) {
		fmt.Println(n)
	}

}

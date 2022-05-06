package main

import (
	"fmt"
	"sync"
	//"time"
)

func producer(num int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < num; i++ {
			out <- i

		}
		close(out)

	}()

	return out
}
func squre(ch <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for k := range ch {
			out <- k * k
		}
		//time.Sleep(time.Second)
		close(out)

	}()

	return out
}
func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	collect := func(in <-chan int) {
		defer wg.Done()
		for k := range in {
			out <- k
		}

	}
	wg.Add(len(ch))
	for _, v := range ch {

		go collect(v)
	}
	go func() {
		wg.Wait()
		close(out)

	}()
	return out

}

func main() {
	in := producer(10)
	ch1 := squre(in)
	ch2 := squre(in)
	ch3 := squre(in)
	out := make(chan int)
	//consumer
	go func() {

		for {
			select {
			case k, ok := <-ch1:
				if !ok {
					ch1 = nil
				} else {
					out <- k
				}

			case k, ok := <-ch1:
				if !ok {
					ch2 = nil
				} else {
					out <- k
				}

			case k, ok := <-ch1:
				if !ok {
					ch3 = nil
				} else {
					out <- k
				}
				if ch1 == nil && ch2 == nil && ch3 == nil {
					return
				}

			}

		}

	}()

	for k := range out {
		fmt.Println(k)
	}
	fmt.Println("end")

}

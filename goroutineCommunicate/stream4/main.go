package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getWaiter() string {
	waiters := []string{
		"Waiter 1",
		"Waiter 2",
		"Waiter 3",
	}
	idx := rand.Intn(len(waiters))
	return waiters[idx]

}

func getChef() string {
	chefs := []string{
		"Chef 1",
		"Chef 2",
		"Chef 3",
	}
	inx := rand.Intn(len(chefs))
	return chefs[inx]

}
func takeOrlder(ordlerId int, wg *sync.WaitGroup, canTakedOrlders chan int, done chan bool) {

	waiter := getWaiter()
	fmt.Printf("%s has taken orlder number %v\n", waiter, ordlerId)
	canTakedOrlders <- ordlerId
	wg.Done()
	select {
	case <-done:
		fmt.Println("case done return")
		return

	}
}

func cookOrlder(wg *sync.WaitGroup, canCookedOrlders chan int, canBringedrlders chan int, done chan bool) {

	for ordlerId := range canCookedOrlders {

		chef := getChef()
		fmt.Printf("%s has brought dishes for orlder number %v\n", chef, ordlerId)
		canBringedrlders <- ordlerId
		wg.Done()
	}

	select {
	case <-done:
		fmt.Println("case done return")
		return

	}

}
func bringOrlder(wg *sync.WaitGroup, canBringedrlders chan int, done chan bool) {

	for ordlerId := range canBringedrlders {
		waiter := getWaiter()
		fmt.Printf("%s has brought dishes for orlder number %v\n", waiter, ordlerId)

		wg.Done()
	}
	select {
	case <-done:
		fmt.Println("case done return")
		return

	}

}
func DealOrlder(orlderId int, wg *sync.WaitGroup, done chan bool) {
	wg.Add(3)

	canCookedOrlders := make(chan int)
	canBringedrlders := make(chan int)
	go takeOrlder(orlderId, wg, canCookedOrlders, done)
	go cookOrlder(wg, canCookedOrlders, canBringedrlders, done)
	go bringOrlder(wg, canBringedrlders, done)

}
func main() {
	start := time.Now()

	var wg sync.WaitGroup
	done := make(chan bool)
	for orlderId := 0; orlderId < 8; orlderId++ {
		DealOrlder(orlderId, &wg, done)
	}

	wg.Wait()
	fmt.Println("wg wait over")
	done <- true
	stop := time.Now()
	fmt.Printf("Time  waste %v \n", stop.Sub(start))
}

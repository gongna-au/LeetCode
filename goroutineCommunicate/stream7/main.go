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
func takeOrlder(ordlerId int, wg *sync.WaitGroup, canTakedOrlders chan int) {

	waiter := getWaiter()
	fmt.Printf("%s has taken orlder number %v\n", waiter, ordlerId)
	canTakedOrlders <- ordlerId
	wg.Done()

}

func cookOrlder(wg *sync.WaitGroup, canCookedOrlders chan int, canBringedrlders chan int) {

	for ordlerId := range canCookedOrlders {

		chef := getChef()
		fmt.Printf("%s is cooked dishes for orlder number %v\n", chef, ordlerId)
		canBringedrlders <- ordlerId
		wg.Done()
	}

}
func bringOrlder(wg *sync.WaitGroup, canBringedrlders chan int) {

	for ordlerId := range canBringedrlders {
		waiter := getWaiter()
		fmt.Printf("%s has brought dishes for orlder number %v\n", waiter, ordlerId)

		wg.Done()
	}

}
func DealOrlder(orlderId int, wg *sync.WaitGroup, wg2 *sync.WaitGroup, canCookedOrlders chan int, canBringedrlders chan int) {
	wg.Add(3)
	go takeOrlder(orlderId, wg, canCookedOrlders)
	go cookOrlder(wg, canCookedOrlders, canBringedrlders)
	go bringOrlder(wg, canBringedrlders)
	wg.Wait()
	wg2.Done()

}
func main() {
	start := time.Now()

	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	canCookedOrlders := make(chan int)
	canBringedrlders := make(chan int)
	for orlderId := 0; orlderId < 8; orlderId++ {
		wg2.Add(1)
		go DealOrlder(orlderId, &wg, &wg2, canCookedOrlders, canBringedrlders)
	}

	wg2.Wait()
	fmt.Println("wg wait over")

	stop := time.Now()
	fmt.Printf("Time  waste %v \n", stop.Sub(start))
}

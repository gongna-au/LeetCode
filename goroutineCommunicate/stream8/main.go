package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

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
func takeOrlder(ordlerId int, canTakedOrlders chan int) {

	waiter := getWaiter()
	fmt.Printf("%s has taken orlder number %v\n", waiter, ordlerId)
	canTakedOrlders <- ordlerId
	wg.Done()
}

func cookOrlder(canCookedOrlders chan int, canBringedrlders chan int) {

	for ordlerId := range canCookedOrlders {

		chef := getChef()
		fmt.Printf("%s is cooked dishes for orlder number %v\n", chef, ordlerId)
		canBringedrlders <- ordlerId
		wg.Done()

	}

}
func bringOrlder(canBringedrlders chan int) {
	for ordlerId := range canBringedrlders {
		waiter := getWaiter()
		fmt.Printf("%s has brought dishes for orlder number %v\n", waiter, ordlerId)
		wg.Done()
	}

}
func DealOrlder(orlderId int, canCookedOrlders chan int, canBringedrlders chan int) {
	wg.Add(3)

	go takeOrlder(orlderId, canCookedOrlders)
	go cookOrlder(canCookedOrlders, canBringedrlders)
	go bringOrlder(canBringedrlders)

}
func main() {
	start := time.Now()

	canCookedOrlders := make(chan int)
	canBringedrlders := make(chan int)
	for orlderId := 0; orlderId < 8; orlderId++ {

		DealOrlder(orlderId, canCookedOrlders, canBringedrlders)
	}

	wg.Wait()
	fmt.Println("wg wait over")

	stop := time.Now()
	fmt.Printf("Time  waste %v \n", stop.Sub(start))
}

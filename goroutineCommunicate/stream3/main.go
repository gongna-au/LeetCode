package main

import (
	"fmt"
	"math/rand"
	"sync"
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
func takeOrlder(ordlerId int, wg *sync.WaitGroup) {

	waiter := getWaiter()
	fmt.Printf("%s has taken orlder number %v\n", waiter, ordlerId)
	wg.Done()
}

func cookOrlder(ordlerId int, wg *sync.WaitGroup) {

	chef := getChef()
	fmt.Printf("%s is cooking orlder number %v\n ", chef, ordlerId)
	wg.Done()
}
func bringOrlder(ordlerId int, wg *sync.WaitGroup) {

	waiter := getWaiter()
	fmt.Printf("%s has brought dishes for orlder number %v\n", waiter, ordlerId)
	wg.Done()

}
func DealOrlder(orlderId int, wg *sync.WaitGroup) {
	wg.Add(3)
	go takeOrlder(orlderId, wg)
	go cookOrlder(orlderId, wg)
	go bringOrlder(orlderId, wg)

}
func main() {
	var wg sync.WaitGroup
	for orlderId := 0; orlderId < 8; orlderId++ {
		DealOrlder(orlderId, &wg)
	}
	wg.Wait()

}

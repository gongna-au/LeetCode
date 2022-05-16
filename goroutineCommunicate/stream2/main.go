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
func takeOrlder(ordlerId int) {
	waiter := getWaiter()
	fmt.Printf("%s has taken orlder number %v\n", waiter, ordlerId)
}

func cookOrlder(ordlerId int) {
	chef := getChef()
	fmt.Printf("%s is cooking orlder number %v\n ", chef, ordlerId)
}
func bringOrlder(ordlerId int) {
	waiter := getWaiter()
	fmt.Printf("%s has brought dishes for orlder number %v\n", waiter, ordlerId)

}
func DealOrlder(orlderId int, wg *sync.WaitGroup) {

	takeOrlder(orlderId)
	cookOrlder(orlderId)
	bringOrlder(orlderId)
	wg.Done()

}
func main() {
	var wg sync.WaitGroup
	for orlderId := 0; orlderId < 8; orlderId++ {
		wg.Add(1)
		go DealOrlder(orlderId, &wg)
	}
	wg.Wait()

}

package main

import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
)

func All(s string, ch chan string) {
	fmt.Println(1)
	ch <- s
}
func Eat(ch chan string) {
	fmt.Println(2)

	s := <-ch
	if s == "eat" {
		fmt.Println("start eating")
		fmt.Println("I am " + s)

	}
}
func Sleep(ch chan string) {
	s := <-ch
	if s == "sleep" {
		fmt.Println("start sleeping")
		fmt.Println("I am " + s)

	}
}
func DeatDouDou(ch chan string) {
	s := <-ch
	if s == "doudou" {
		fmt.Println("start sleeping")
		fmt.Println("I am " + s)

	}
}
func main() {
	ch := make(chan string, 1)
	go Eat(ch)
	All("eat", ch)

}

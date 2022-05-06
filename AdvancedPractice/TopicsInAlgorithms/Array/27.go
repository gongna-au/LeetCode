package main

import (
	"fmt"
	//"runtime"
	//"sync"
	//"time"
)

func Eat() chan string {
	ch := make(chan string, 1)
	ch <- " start eating"

	return ch

}
func Sleep() chan string {
	ch := make(chan string, 1)
	ch <- " start sleeping"
	return ch

}
func DeatDouDou() chan string {
	ch := make(chan string, 1)
	ch <- " start do Dou dou"
	return ch

}
func main() {
	ch1 := Eat()
	ch2 := Sleep()
	ch3 := DeatDouDou()

	select {
	case s := <-ch1:
		fmt.Println(s)
		fmt.Println("I am eating")
	case t := <-ch2:
		fmt.Println(t)
		fmt.Println("I am sleeping")
	case y := <-ch3:
		fmt.Println(y)
		fmt.Println("I am do Dou dou ")
	}

}

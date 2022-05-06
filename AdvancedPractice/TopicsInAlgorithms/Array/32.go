package main

import "fmt"

func test1() int {
	fmt.Println("In test1")
	return 9

}
func test2() int {
	fmt.Println("In test2")
	return 8

}
func All(a int, b int, c int) int {
	fmt.Println("In All")
	return a + b + c

}

func main() {
	fmt.Println(All(test1(), test2(), 1))
}

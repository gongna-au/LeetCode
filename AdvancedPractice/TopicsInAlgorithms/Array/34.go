package main

import (
	"fmt"
	"time"
)

func test(s string) {
	for i := 0; i < 4; i++ {
		fmt.Println("i=", i, s)
	}
}
func main() {
	test("bird")
	go test("tree")
	go func(s string) {

		fmt.Println(s)

	}("plant")
	test("person")
	time.Sleep(time.Second)
}

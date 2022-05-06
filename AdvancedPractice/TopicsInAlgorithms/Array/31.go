package main

import "fmt"

func test() (result int) {
	//result=9
	defer func(temp int) { //result=0
		fmt.Println("result", result)
		fmt.Println("temp", temp)

		temp = result + temp //sum =5
		fmt.Println("temp", temp)
	}(result)

	return 9

}

func main() {
	fmt.Println(test())
}

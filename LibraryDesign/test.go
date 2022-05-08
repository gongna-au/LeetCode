package main

import "fmt"

func main() {
	array := make(map[string]interface{})
	array["test"] = nil
	k, ok := array["test"]
	if ok {
		if k != nil {
			fmt.Println(k)
		} else {
			fmt.Println("k is nil")
		}

	} else {
		fmt.Println("a")
	}

}

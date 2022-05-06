package main

import "fmt"
func isPowerOfTwo(n int ) bool {
	if   n > 0 && n & (n-1) == 0{
		return true 
	}else{
		return false
	}


}
func main(){
	n := 8
	m := 9
	fmt.Println(isPowerOfTwo(n))
	fmt.Println(isPowerOfTwo(m))

}
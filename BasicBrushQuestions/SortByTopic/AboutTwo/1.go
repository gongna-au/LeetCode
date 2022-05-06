package main
import "fmt"
func AboutTwo( n int) bool{
	if n !=1 && n%2 != 0 {
		return false
	}
	if n == 1 {
		return  true 
	}else {
		return AboutTwo(n/2)
	}
	

}
func main(){
	n := 8
	m := 9
	fmt.Println(AboutTwo(n))
	fmt.Println(AboutTwo(m))

}


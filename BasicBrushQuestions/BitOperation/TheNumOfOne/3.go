package main

import "fmt"
func NumOfOne( n int)(result int){
	//一个数与比自己小1的数做位运算就是把自己最低位的1变为0
	if n==0{
		return 0
	}else{
		return 1+NumOfOne(n&(n-1))
	}
	
}

func  main()  {
	n := 0b00000000000000000000000000001011
	m := 0b00000000000000000000000010000000
	p := 0b11111111111111111111111111111101
	fmt.Println(NumOfOne(n))
	fmt.Println(NumOfOne(m))
	fmt.Println(NumOfOne(p))
}

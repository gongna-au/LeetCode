package main
import "fmt"
func DigitalRotationMatrix(root int,k int,array [][]int) ( x int  , y int  ){
	
	
		if root == 1 {
			x =0
			y =0

		}

		for i:=0 ;i < k;i++{
			array[x][y] = root+i
			x++
		}

		//i= 0 j=5
		
		for i :=0 ;i < k; i++{
			array[x][y] = root+ k+i
			y++
		}

	    // i=5 j=5
		
		for i :=0 ;i < k;i++{
			array[x][y] = root+2*k+i
			x--
			
		}
		//j=0 i=5

		for i:=0 ;i < k;i++{

			array[x][y] = root+ 3*k+i
			y--
		}
		
	return DigitalRotationMatrix(root+k*4,k-2,array)
	
}
func main()  {
	k := 5

    array := make([][]int, k+1) // 二维切片，3行
    for i := range array {
        array[i] = make([]int, k+1) // 每一行4列
    }
   
	DigitalRotationMatrix(1,k,array)
		for _,v := range array{

			for _,v2 := range v{
				fmt.Printf("%v  ",v2)
	
			}
			fmt.Println()
		}
}


/* 
	0  1  2  3   4   5
0	1 20 19 18  17  16
1	2 21 32 31  30  15
2	3 22 33 36  29  14
3	4 23 34 35  28  13
4	5 24 25 26  27  12
5	6  7  8  9  10  11 
 */
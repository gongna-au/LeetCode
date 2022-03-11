package main
import (
	"fmt";
)
// 用分治法求数组的最大子段和问题

func maxSubsectionAnd(i int ,j int, array []int ) (subi int ,subj int ,max int){
	if i==j {
		if array[i]>=0{
			return i ,j
		}else{
			return 
		}
	
	}else if i < j{
		for i<j{
			if array[i]0{

			}
			i++
		}
		
	}

	

	lefti,leftj := maxSubsectionAnd(i,mid,array)
	righti,rightj := maxSubsectionAnd(mid+1,j,array)
	crosingi,crosingj := maxSubsectionCrosing(i,j,mid)
i
	return max(leftmax,rightmax,1)
}

func maxSubsectionCrosing(i int ，j int , array []int){
	sumLeft:=0
	for  k:= mid ;k >=i ;k-- {
		sumLeft = sumLeft+array[k]

	}







}
func max(a int ,b int ,c int) int {
	max:=a
	if b >= max {
		max=b
	}

	if c >= max {
		max=c
	}

	return max
}



func main(){
	fmt.Println(5/2)
	// 0 1 2 3 4 


}



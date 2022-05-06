package main
import "fmt"
func intersection(a []int ,b []int ) []int{
	temp:=map[int]int{}
	for _,v:=range  a{
		temp[v]+=1
	}
	k:=0
	for _,v:=range b{
		if temp[v]>=1{
			temp[v]-=1
			b[k]=v 
			k++
		}
	}
	return b[0:k]

}
func main(){
	a:=[]int{2,3,2,4,5}
	b:=[]int {2,2}
	c:=intersection(a,b)
	fmt.Println(c)

}

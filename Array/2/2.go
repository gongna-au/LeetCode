package main

import "fmt"

type Arrayslice []int

func(a Arrayslice ) len() int{
	return len(a)
}
func min(a Arrayslice,b Arrayslice ) int {
	if (a.len()<=b.len()){
		return a.len()
	}else{
		return b.len()
	}
}

func intersection(a Arrayslice,b Arrayslice  ){
	blenth:=b.len()
	temp:=min(a,b)
	
	var tempslice []int
	
	k :=0
	for i:=0;i<temp;i++{
		
		 for j:=0;j<blenth;j++{
			 if(a[i]==b[j]){
				tempslice=append(tempslice, a[i])
				k++
				break
			}
		}
	}
	fmt.Println(k)
	for i:=0;i<k;i++{
		fmt.Println(tempslice[i])
	}
	
	



}





func main(){

	 var a=[]int{1,2,3,4}
	 var b=[]int{3,2,3,4}
	 intersection(a,b)





}
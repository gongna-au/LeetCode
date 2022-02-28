package main

import "fmt"
type  Arrayslice []int
func ( a Arrayslice) len() int {
	return len(a)
}

func( a Arrayslice) less( x int ,y int) bool{
	return a[x]<=a[y]
}

func (a Arrayslice) swap( x int , y int ){
	
	if(a.less(x,y)){
		
		a[x],a[y]=a[y],a[x]

	}
	//fmt.Println(a[x])
	//fmt.Println(a[y])


}
func  Swap(a Arrayslice){
	len:=a.len()
	 
	for i:=0; i<len-1; i++ {
		for j:=0;j<len-1;j++{
			if a.less(j,j+1) {
				a.swap(j,j+1)
			}
		}

	}
	for i:=0;i<len;i++{
		fmt.Println(a[i])
	}
}





/*func  main(){
	var array=[]int {1,2,3,4}
	Swap( array)

}*/





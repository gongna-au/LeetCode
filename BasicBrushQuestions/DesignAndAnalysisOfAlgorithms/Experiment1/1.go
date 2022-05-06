package main
import "fmt"
//移除相同元素并保持元素顺序不变
func DeleteSameItem(w []int)[]int{
	if len(w)==0{
		return []int{ }
	}
	result:=[]int{}
	wMap:=make(map[int]int,len(w))
	for _,v:=range w{
		 wMap[v]=wMap[v]+1

	}

	for k, _ := range wMap {
		result=append(result,k)
		
	}
	return result

}
func main(){
	w:=[]int{1,2,2,3,3,4,5,6,7}
	fmt.Println(DeleteSameItem(w))
}


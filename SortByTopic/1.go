//插入排序
package main
import "fmt";
func InsertSort(array []int)([]int){
	for i:=1; i < len(array)-1 ; i++ {
		
		for j:=i; j >0; j--{
			if array[j-1] > array[j] {
				array[j] ,array[j-1] = array[j-1],array[j]
			}

		}
		

	}
	return array
}
func main(){
	array := []int{2,87,0,5,8,53,83,9,45}
	fmt.Println(InsertSort(array))
	
}


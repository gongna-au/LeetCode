//给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
//给定 nums = [3,2,2,3], val = 3,
//2 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//3 你不需要考虑数组中超出新长度后面的元素。

package main
import "fmt"

func deleteInPlace(a []int, val int )int {
	result:=[]int{}
	for _,v:=range a{
		if v!=val{
			result=append(result,v)
		}
	} 
	fmt.Println(result)
	 return len(result)

}
func main(){
	nums := []int {3,2,2,3}
	val := 3
	fmt.Println(deleteInPlace(nums ,val))
}



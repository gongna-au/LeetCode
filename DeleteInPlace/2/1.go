
//给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
//给定 nums = [3,2,2,3], val = 3,
//2 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//3 你不需要考虑数组中超出新长度后面的元素。
// 原地删除要求你不要使用额外的数组空间


package main
import "fmt"

func deleteInPlace(a []int, val int )int {

	
	for k:=0;k<len(a);{
		if a[k]==val{
			a=append(a[:k],a[k+1:]...)
			//此时k上放了新的元素，我们依旧需要判断k上放的这个元素和所给的值相不相等

		}else{
			k++
		}

	}
	fmt.Println(a)
	return len(a)
	
}
func main(){
	nums := []int {3,2,2,3}
	val := 3
	fmt.Println(deleteInPlace(nums ,val))
}



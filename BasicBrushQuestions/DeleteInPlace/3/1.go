//删除排序数组中的重复项
//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次。
//返回移除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1)额外空间的条件下完成
//给定数组 nums = [1,1,2],
//2 函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//3 你不需要考虑数组中超出新长度后面的元素。
//示例 2:
//1 给定 nums = [0,0,1,1,1,2,2,3,3,4],
//2 函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
//3 你不需要考虑数组中超出新长度后面的元素。
package main
import "fmt"
func deleteSameItem(a []int) []int {
	i:=0
	j:=1
	for ; i<len(a)-2; {
		if(a[j]==a[i]){
			a=append(a[:j],a[j+1:]...)
		}else{
			j++
			i=j-1
		}
	}
	fmt.Println(len(a))
	return a

	
}
func main()  {
	nums := []int{0,0,1,1,1,2,2,3,3,4}

	fmt.Println(deleteSameItem(nums))

	

	
}
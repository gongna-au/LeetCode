package main
import "fmt"
/*两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，
并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
示例:
1 给定 nums = [2, 7, 11, 15], target = 9
2 因为 nums[0] + nums[1] = 2 + 7 = 9
3 所以返回 [0, 1]
*/

func SumOfTwo(a []int ,target int) []int {
	
	m:=make(map[int]int)
	for k,v:=range a{
		m[v]=k
	}
	for _,v:=range a{
		if k,ok := m[target-v];ok{
			 return []int{m[v],k}
		}

	}
	
	
	fmt.Println("Can not find these two number")
	return []int{}
}
func main()  {
	a:=[]int {2, 3, 6, 15}
	target :=9
	fmt.Println(SumOfTwo(a,target))
}
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//示例 1:
//1 输入: [1,2,3,4,5,6,7] 和 k = 3
//2 输出: [5,6,7,1,2,3,4]
//要求使用空间复杂度为 O(1) 的 原地 算法。
package main
import ("fmt";
		
)




func runOneStep(a []int) []int{
	lenth:=len(a)
	temp:=a[lenth-1]
	for i:=lenth-1;i>=1;i--{
		a[i]=a[i-1]
	}
	a[0]=temp
	return a
}

func rotateArray( a []int ,k int )[]int{
	
	for i:=0; i<k; i++{
		a=runOneStep(a)
	}
	
	return a

}
func main(){
	k:=2
	a:= []int {-1,-100,3,99}
	fmt.Println(rotateArray(a,k))

}
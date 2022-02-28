//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//示例 1:
//1 输入: [1,2,3,4,5,6,7] 和 k = 3
//2 输出: [5,6,7,1,2,3,4]
package main
import ("fmt";
		"math"
)





func rotateArray( a []int ,k int )[]int{
	alenth:=len(a)
	
	a = append(a, a...)
	result:=[]int{}
	step:=int(math.Abs(float64(alenth-k)))
	j:=step
	result=append(result,a[j:alenth+step]...)
	
	return result

}
func main(){
	k:=2
	a:= []int {-1,-100,3,99}
	fmt.Println(rotateArray(a,k))
}
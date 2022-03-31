package main
/* 缺失数字
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
示例 1:1 输入: [3,0,1]
2 输出: 2
示例 2:
1 输入: [9,6,4,2,3,5,7,0,1]
2 输出: 8
说明:
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现? */
import "fmt"

func TheNumberLost(array []int )( int){
	sum1:=0
	max:=0
	for i:=0; i<len(array); i++{
		sum1= sum1+array[i]
		if array[i] > max{
			max=array[i]
		}

	}
	sum2:=0
	for  i:=0;i<= max ;i++{
		sum2=sum2+i
	}
	return sum2-sum1

}
func main(){
	array := []int {9,6,4,2,3,5,7,0,1}
	fmt.Println(TheNumberLost(array))

}
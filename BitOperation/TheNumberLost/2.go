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
//使用异或解决
//在循环的过程中把出现了两次的值消掉

func TheNumberLost(array []int )(result int){
	
	max:=0
	
	for i:=0; i<len(array); i++{
		
		if array[i] > max{
			max=array[i]
		}

	}
	
	for  i:=0;i <= max ;i++{
		if i<len(array){
			result = result ^ array[i]^i

		}else{
			result = result^i

		}
			
	}
	return result

}
func main(){
	array := []int {9,6,4,2,3,5,7,0,1}
	fmt.Println(TheNumberLost(array))

}
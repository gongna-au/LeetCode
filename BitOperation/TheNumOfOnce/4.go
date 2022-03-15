/* 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只
出现了一次的元素。说明：你的算法应该具有线性时间复杂度。你可以不使用额外空间来实现吗？ */
package main

import "fmt"
func NumberAttendOnce(array []int )( result int){
	// a?a?a=0
	result =0
	
	for i:=0;i <32;i++{
		//保证每一个位的初始都是从零开始相加
		sum:=0
		for j:=0;j<len(array);j++{
			sum = sum+(( array[j]>>i ) & 1)
		}
		//用到了或运算是把每一个位置的1填到相应的位置上
		result =result| (sum%3)<<i
	}
	return result

	// 101011
	// 101011
	// 101011
	// 100001
	
}
func main(){
	array1 := []int {2,2,1,2,3,3,3,4,4,4}
	array2 := []int {5,2,2,2,3,3,3}
	fmt.Println(NumberAttendOnce(array1))
	fmt.Println(NumberAttendOnce(array2))


}
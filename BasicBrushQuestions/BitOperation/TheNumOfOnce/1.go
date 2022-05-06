/* 只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出
现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例 1:
1 输入: [2,2,1]
2 输出: 1
示例 2:1 输入: [4,1,2,1,2]
2输出: 4 

*/
package main

import "fmt"
func NumberAttendOnce(array []int )(int){
	mapList := map[int]int{}
	for i:=0;i<len(array)-1;i++{
		mapList[array[i]]= mapList[array[i]]+1
	}
	i:=0
	for ;i<len(array)-1;i++{
		 if mapList[array[i]] !=2{
			 break
			 

		 }
	}
	return array[i]



}
func main(){
	array1 := []int {2,2,1}
	array2 := []int {4,1,2,1,2}
	fmt.Println(NumberAttendOnce(array1))
	fmt.Println(NumberAttendOnce(array2))

}
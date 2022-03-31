/* 最小的k个数
今天分享一道比较简单的题目，希望大家可以5分钟掌握！
01、题目示例
最小的k个数
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则
最小的4个数字是1、2、3、4。
示例 1：
1 输入：arr = [3,2,1], k = 2
2 输出：[1,2] 或者 [2,1]
*/
package main

import (
	"fmt"
	"os"
	"sort"
)

func MinKNumber(array []int, k int) (result []int) {
	if len(array) <= 4 {
		os.Exit(-1)
	}
	temp := array[:len(array)-1]
	startArray := GetStartArray(temp, k)
	fmt.Println(array)
	for _, v := range array {

		startArray = AdjustOrder(v, startArray)

	}
	return startArray

}

func GetStartArray(array []int, k int) []int {
	temp := array
	sort.Ints(temp)
	return temp[:k]
}
func AdjustOrder(value int, array []int) []int {
	return Adjust(0, len(array)-1, array, value)

}

func Adjust(start int, end int, array []int, value int) []int {
	if start > end {
		//2 3
		temp := array[start : len(array)-1]
		array = append(array[:start], value)
		array = append(array, temp...)
		return array

	}
	if value <= array[start] || value >= array[end] {
		return array

	} else {
		array = Adjust(start+1, end-1, array, value)

	}
	return array

}
func MinKNum(array []int, k int) []int {
	sort.Ints(array)
	return array[:k]
}
func main() {
	array := []int{1, 6, 3, 2, 0, 7, 15, 18, 20, 34}

	fmt.Println(MinKNum(array, 4))
	// 1 2 3 6

}

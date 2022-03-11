/* 
滑动窗口最大值
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。
滑动窗口每次只向右移动一位。返回滑动窗口中的最大值。
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到
在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值所构成的数组 
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:
滑动窗口的位置
最大值
--
1  3  -10  -3  -9  5  3  6  7

1  3  -1  -3  5  3  6  7

1  3  -1  -3  5  3  6  7

1  3  -1  -3  5  3  6  7

1  3  -1  -3  5  3  6  7

1  3  -1  -3  5  3  6  7

*/
/* 
1

3

3 -10
 
3 -10 -3

3 -10 -3 -9

5

5 3
 
6

7


 */
package main

import (

	"fmt"
	//"math"
	//"sort"
)

func SlideWindowSeries(array []int, n int ) (result []int) {
	variableSlice := []int {}
	lastmax := 0
	
	for i := 0;i < len(array) ; i++ {
		if len(variableSlice) >= 3 {
			// 
			variableSlice = variableSlice[ len(variableSlice)-1-(n-1) +1:]
			lastmax = max(variableSlice[0],variableSlice[1])

		}
		if len(variableSlice) == 0 {
			variableSlice = append(variableSlice,array[i])
			lastmax = array[i]
		}
		
		if  array[i] < lastmax && len(variableSlice) <=2  {
			variableSlice = append(variableSlice,array[i])
		}
		if array[i] > lastmax {
			variableSlice = []int {}
			variableSlice = append(variableSlice,array[i])
			lastmax = array[i] 
		}
		result = append(result,lastmax)
	
	}
	return result

}
func max(a int, b int  )int {
	if a > b {
		return a
	}else {
		return b
	}

}
/* 
5         		5
5 4       		5
    6     		6
	  7   		7
	  7 6 		7
	  7 6 5 	7
	    6 5 4   6 
		     7  7
	   */

func main()  {
	array := []int {1,3,-1,-3,5,3,6,7}
	/* 1
	   3 
	   3
	   3
	   5
	   5
	   6
	   7 
	   */
	k:=3
	fmt.Println(SlideWindowSeries(array,k))
	
}



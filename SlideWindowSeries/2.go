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
 )
 
 func SlideWindowSeries(array []int, n int ) []int {
	 if len(array)==0 {
		return []int{} 
	 }
	 result := []int{}
	 queue := []int {}
	
	 for i:=0; i<=len(array)-k; i++ {
		 if i==0{
			queue := append(queue,array[0])
		 }
		 if i>0 {
			if array[i] > queue[0] {
				queue[0] = array[i]
			}else{

				

			}


		 }




		result := append(result, queue[0])
		
	 }
	 return result
 
 }
 func main()  {
	 array := []int {1,3,-1,-3,5,3,6,7}
	 k:=3
	 // 9 6 7 8 5 4 3 2 1
	 /* 9
	 	9 6
		9 6 7 
		9 6 7 8
		9 8 7 6 5
		0 1 2 3 4 5 
	  */
	


	 fmt.Println(SlideWindowSeries(array,k))
	 
 }
 
 
 
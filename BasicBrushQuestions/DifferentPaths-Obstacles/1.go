/* 不同路径 - 障碍物
一个机器人位于一个 m x n 网格的左上角，起始点在下图中标记为“Start”。机器人每次只能向下或
者向右移动一步。机器人试图达到网格的右下角，在下图中标记为“Finish”。现在考虑网格中有障
碍物。那么从左上角到右下角将会有多少条不同的路径？ 问总共有多少条不同的路径？ */
/* 输入:
 [
  [0,0,0],
  [0,1,0],
  [0,0,0]
  ]
   输出: 2


解释:
 3x3 网格的正中间有一个障碍物。
 从左上角到右下角一共有 2 条不同的路径：
 \1. 向右 -> 向右 -> 向下 -> 向下
 \2. 向下 -> 向下 -> 向右 -> 向右
*/
package main

import "fmt"

func DifferentPath(array [][]int) int {
	dp := make([][]int, len(array))
	for i := 0; i < len(array[0]); i++ {
		dp[i] = make([]int, len(array[0]))
	}
	dp[0][0] = 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			if i == 0 && j == 0 {
				continue
			} else if array[i][j] == 1 {
				dp[i][j] = 0

			} else if i == 0 {
				dp[i][j] = 1
			} else if j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}
	return dp[len(array)-1][len(array[0])-1]
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	array := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(DifferentPath(array))

}

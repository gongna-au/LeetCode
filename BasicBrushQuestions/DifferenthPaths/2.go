/*
leetcode-63
⼀个机器⼈位于⼀个 m x n ⽹格的左上⻆ （起始点在下图中标记为“Start” ）。
机器⼈每次只能向下或者向右移动⼀步。
机器⼈试图达到⽹格的右下⻆（在下图中标记为“Finish”）。
现在考虑⽹格中有障碍物。那么从左上⻆到右下⻆将会有多少条不同的路径
*/
package main

import (
	"fmt"
)

func differentPath(dp [][]int) int {

	//初始化初始状态
	for i := 0; i < len(dp[0]); i++ {
		dp[0][i] = 1
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
	}

	//设置障碍物
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if i == 0 || j == 0 {
				continue

			} else if dp[i][j] == -1 {
				dp[i][j] = 0
				continue
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}
	return dp[len(dp)-1][len(dp[0])-1]

}

/*
0  0  0  0
0 -1  0  0
0  0  0  0
0 -1 -1  0
*/
func main() {
	dp := [][]int{
		{0, 0, 0, 0},
		{0, -1, 0, 0},
		{0, 0, 0, 0},
		{0, -1, -1, 0},
	}
	fmt.Println(differentPath(dp))

}

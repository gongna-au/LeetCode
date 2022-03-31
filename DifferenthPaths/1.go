/*
leetcode-62
一个机器人位于一个 m x n 网格的左上角，起始点在下图中标记为“Start”。机器人每次只能向下或
者向右移动一步。机器人试图达到网格的右下角，在下图中标记为“Finish”。 问：总共有多少条不
同的路径？ */

/*
输入: n = 2 m = 3
输出: 3
解释:
从左上角开始，总共有 3 条路径可以到达右下角。
\1. 向右 -> 向右 -> 向下
\2. 向右 -> 向下 -> 向右
\3. 向下 -> 向右 -> 向右


示例 2:
输入: m = 7, n = 3
输出: 28
*/
package main

import "fmt"

/*
1 1 1
1 0 0


x 0->1 0-1-1-1
y 0->2 0-0-1-2
x 0->1 0-0-0-1
y 0->2 0-1-2-2
 	0-1-2-3

1 1 1 1 1 1 1 1
1
1




*/
func DifferentPaths(row int, list int) int {

	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, list)
	}
	//dp[i][j] = dp[i][j-1]   ||  dp[i-1][j]
	for i := 0; i < list; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < row; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < row; i++ {
		for j := 0; j < list; j++ {
			//状态转移的点
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]

		}
	}
	return dp[row-1][list-1]

}

func main() {
	row := 3
	list := 7
	fmt.Println(DifferentPaths(row, list))

}

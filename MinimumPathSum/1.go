/* leetcode-64
 */
package main

import "fmt"

/*
Input:
[
[1,3,1],
[1,5,1],
[4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
给定⼀个包含⾮负整数的 m x n ⽹格，
请找出⼀条从左上⻆到右下⻆的路径，
使得路径上的数字总和为最⼩。
说明：每次只能向下或者向右移动⼀步。 */
//原地dp
func MinimumPathSum(dp [][]int) int {
	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + dp[i][0]
	}

	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = dp[0][i-1] + dp[0][i]
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if i == 0 || j == 0 {
				continue
			} else {
				dp[i][j] = dp[i][j] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func main() {
	dp := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(MinimumPathSum(dp))

}

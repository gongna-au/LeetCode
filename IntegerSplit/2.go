/*
第343题：整数拆分
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。返回你可以获
得的最大乘积。
示例 1:
  输入: 2
  输出: 1
  解释: 2 = 1
  1 × 1 = 1。
示例 2:
  输入: 10
  输出: 36
  解释: 10 = 3
  3 × 3 × 4 = 36

  /

说明: 你可以假设 n 不小于 2 且不大于 58。 */
//用动态规划做

//dp[i]=max(dp[i],(i-j)*max(dp[j],j))
package main

import (
	"fmt"
)

func IntegerSplit(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 2; i <= n; i++ {
		for j := 2; j < i; j++ {
			// 2 不分解 7 2*3 3*4
			// 8 2*6 3*5 4*4
			// 9 2*7 3*6  4*5
			// max()里面的dp[i]代表的是前一次得到的dp的状态
			// 前一次状态和
			dp[i] = max(dp[i], j*(i-j), j*dp[i-j])
		}
	}
	return dp[n]

}
func max(a int, b int, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func main() {
	fmt.Println(IntegerSplit(10))

}

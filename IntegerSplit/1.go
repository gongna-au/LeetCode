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

//dp[i]=max(dp[i],(i-j)*max(dp[j],j))
package main

import (
	"fmt"
	//"sort"
)

func IntegerSplit(n int) int {
	if n <= 0 {
		return 1
	}

	if n%3 == 0 {

		if n-3 > 0 {
			return IntegerSplit(n-3) * 3

		} else {
			return 1 * 3

		}
	} else if n%3 == 2 {
		return 2 * IntegerSplit(n-2)
	} else {

		return 2 * IntegerSplit(n-2)

	}

}

func main() {
	fmt.Println(IntegerSplit(8))

}

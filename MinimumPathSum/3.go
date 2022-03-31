package main

import "fmt"

//压缩空间进行dp
func MinimumPathSum(dp [][]int) int {
	temp := make([]int, len(dp[0]))

	//InitfirstRow
	for i := 0; i < len(dp); i++ {
		temp[i] = dp[0][i]
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			temp[j] = temp[j-1]

		}

	}
	return temp[len(dp[0])-1]
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

/*
1 1 1 1      1 1 1 1
1 1 1 1      1
*/

package main

import "fmt"

//压缩空间进行dp
func MinimumPathSum(dp [][]int) int {
	temp := make([]int, len(dp[0]))

	//InitfirstRow
	//因为第一行只能依赖左边
	temp[0] = dp[0][0]
	for i := 1; i < len(dp); i++ {
		temp[i] = temp[i-1] + dp[0][i]
	}
	//第二行左边的第一个只能依赖上面的

	for i := 1; i < len(dp); i++ {
		for j := i; j < len(dp[0]); j++ {
			if j == 0 {
				//对于第一列的数据只依赖于上面一行的第一列
				temp[0] = temp[0] + dp[i][0]
			} else {
				//temp[j]是上一行的数据
				temp[j] = min((temp[j-1]), temp[j]) + dp[i][j]

			}

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

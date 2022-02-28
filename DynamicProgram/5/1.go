
/*
三角形最小路径和
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
例如，给定三角形：
    2                            0
   3 4                          1 2
  6 5 7                        3 4 5
 4 1 8 3                      6 7 8 9

]
0
1 2
3 4 5
6 7 8 9
则自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）*/

package main
import (
	"fmt";
	"sort"
)

func MinimumPathOfTriangular(a [][]int ) int {
	
	var dp [4][4]int
	dp[0][0]=a[0][0]	
	for i:=1;i<4;i++{
	
		for j:=0;j<i+1;j++{
			if j==0 {
				dp[i][j]=a[i][j]+dp[i-1][j]
				
			}else if j==i{
				dp[i][j]=a[i][j]+dp[i-1][j-1]
				
			}else{
				dp[i][j]=a[i][j]+min(dp[i-1][j],dp[i-1][j-1])

			}
			//fmt.Println(dp[i][j])
			
		}
				
	}
	result:=[]int{dp[3][0],dp[3][1],dp[3][2],dp[3][3]}
	fmt.Println(result)
	sort.Ints(result)
	 return result[0]
}
func min(a int,b int)int{
	if a>=b{
		return b
	}else{
		return a
	}
}
func main(){
	a := make([][]int, 4)
	a =[][]int {{2},{3,4},{6,5,7},{4,1,8,3}}
	fmt.Println(MinimumPathOfTriangular(a))

	




}
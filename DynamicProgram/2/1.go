/*
最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回
其最大和。
示例:
1 输入: [-2,1,-3,4,-1,2,1,-5,4],
2 输出: 6
3 解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/ 
package main
import (
	"fmt";
	"sort"
)
func maximumSuborderSum( a []int) int  {
	dp:=make([]int,len(a))
	for i:=0;i<len(a);i++{
		
		if i==0{
			dp[i]=a[i]
		}else{
			if dp[i-1]>0{
				dp[i]=dp[i-1]+a[i]
			}else{
				dp[i]=a[i]
			}
			
		}
	
	}
	sort.Ints(dp)
	return dp[len(dp)-1]	
}
func main(){

	a:=[]int {-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maximumSuborderSum(a))

}
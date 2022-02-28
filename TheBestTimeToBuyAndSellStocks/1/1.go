// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//输入: [7,1,5,3,6,4]
//输出: 7
//解释： 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易
//所能获得利润 = 5 - 1 = 4 。 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格
//= 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3
package main
import (
	"fmt";
	"sort"
)



func  MaximumProfit(a []int) int{
	
	
	i:=0
	j:=1
	k:=0
	f:=0
	b:=[]int{}
	
	for k<len(a)-1{
		i=k
		j=k+1
		
		maxProfit:=0
		for f=k+1;j<len(a);f++ {
			if(a[f]>a[i]){
				maxProfit=maxProfit+(a[f]-a[i])
				i=f+1
				j=i
			}
			j++
		}
		b=append(b,maxProfit)
		k++
	}
	sort.Ints(b)
	return b[len(b)-1]
	
}
func main(){
	a :=[]int {1,2,3,4,5}
	maxProfit:=MaximumProfit(a)
	fmt.Println(maxProfit)
}
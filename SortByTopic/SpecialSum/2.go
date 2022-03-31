/* 连续n个数的和
求 1 2 ... n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句
（A?B:C）。
示例 1：
1
输入: n = 3 输出: 6
示例 2：
1
输入: n = 9 输出: 45 */
package main


import "fmt"
func Plus(n *int ,m int)bool{
	*n = *n+m
	return true
}
//如果 不使用if 和递归的解法
func SpecialSum(n int)int {
	_ = n >0 && Plus(&n,SpecialSum(n-1))
	return n
}

func main(){

	n := 9
	fmt.Println(SpecialSum(n))

}
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
//如果使用if 和递归的解法
func SpecialSum(n int)int {
	if n==0{
		return 0
	}else{
		return n+SpecialSum(n-1)
	}





}
func main(){

	n := 9
	fmt.Println(SpecialSum(n))

}
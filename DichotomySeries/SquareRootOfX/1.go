/* 
x的平方根（69）
今天继续为大家分享二分法系列篇的内容，看一道比较简单的题目。
题目分析
这道题目是比较简单，但我认为同时也是非常经典，建议大家掌握！
x的平方根
计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数
部分将被舍去。

*/
package main

import "fmt"
func SquareRootOf(x float64)( result float64 ){
	if x == 1.0{
		return 1
	}
	low := 1.0
	hight := x
	mid := (hight+low)/2.0
	//定义精确度

	accrucy := 0.00001

	for {
		if mid*mid > x {
			hight=mid
			mid = (hight+low)/2.0
		}else if mid*mid < x{
			if x-mid*mid < accrucy{
				return mid
			}else{
				low=mid
				mid = (hight+low)/2.0

			}

		}else  {
			return mid
		}
	}
}
func main(){
	var x float64
	var y float64
	x = 9
	y = 8
	fmt.Println(SquareRootOf(x))
	fmt.Println(SquareRootOf(y))

}
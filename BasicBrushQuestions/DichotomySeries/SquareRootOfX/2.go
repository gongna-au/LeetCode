/* 
x的平方根（69）
今天继续为大家分享二分法系列篇的内容，看一道比较简单的题目。
题目分析
这道题目是比较简单，但我认为同时也是非常经典，建议大家掌握！
x的平方根
计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数
部分将被舍去。

*/
//用递归法解
package main

import "fmt"
func SquareRootOf(low float64, high float64, target float64)( result float64 ){
	accrucy := 0.000001
	mid := (low+high)/2.0
	if target  == mid*mid{
		return mid
	}else if mid*mid < target{

		if (target-mid*mid) < accrucy {
			return mid
		}else{

			return SquareRootOf(mid,high,target)
		}
		

	}else if (mid*mid) > target{
		return SquareRootOf(low,mid,target)

	}else{
		return mid
	}
		
}
func main(){
	var x float64
	var y float64
	x = 9
	y = 8

	fmt.Println(SquareRootOf(1,x,x))
	fmt.Println(SquareRootOf(1,y,y))

}
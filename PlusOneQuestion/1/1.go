package main
import "fmt"
//plusOneQuestion
/*加一
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，
这个整数不会以零开头。
示例 1:
1 输入: [1,2,3]
2 输出: [1,2,4]
3 解释: 输入数组表示数字 123。
示例 2:
1 输入: [4,3,2,1]
2 输出: [4,3,2,2]
3 解释: 输入数组表示数字 4321
2020213760 龚娜 第一次作业
1 输入: [9,9,9,9]
2 输出: [1,0,0,0,0]

*/

//为什么递归解不出来？
func plusOne(a []int ) int[]{
	benchmark=len(a)-1
	 if a[benchmark]< 9 {
		 a[benchmark]=a[benchmark]+1
		 return  a
	 }else{
		plusOne(a[0:benchmark])

	 }
}

func main()  {
	a:=[]int {1,2,3}

	fmt.Println(plusOne(a))




	
}
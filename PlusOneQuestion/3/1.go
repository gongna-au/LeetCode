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

1 输入: [9,9,9,9]
2 输出: [1,0,0,0,0]

*/

//为什么递归解不出来？
func PlusOne(a []int)[]int {
	// carry为进位
	carry:=0
	benchmark:=len(a)-1
	for i:=benchmark;i>=0;i--{

		if i==benchmark {
			a[i]=a[i]+1
		}
		
		if carry==1{
			a[i]=a[i]+1
		}

		if a[i]== 10{
			a[i]=a[i]%10
			carry=1
		}else{
			carry=0
		}


	}
	if a[0]==0{
		b:=[]int {1}
		b=append(b,a...)
		return b

	}else{
		return a
	}
}

func main()  {
	a:=[]int {9,9,8,9}

	fmt.Println(PlusOne(a))

}
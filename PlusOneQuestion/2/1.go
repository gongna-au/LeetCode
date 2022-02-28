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

	benchmark:=len(a)-1
	result:=[]int {1}
	j:=benchmark
	for ;j>=0;j--{
		if a[j]<9{
			a[j]=a[j]+1
			return a
		}else{
			a[j]=0
			
		}
		
	}
	if j==-1{
		result=append(result,a...)
		return result

	}
	return a

}

func main()  {
	a:=[]int {9,9,8,9}

	fmt.Println(PlusOne(a))

}
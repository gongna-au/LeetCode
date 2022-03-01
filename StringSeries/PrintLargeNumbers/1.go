// 题目：大数打印
// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
// 输入: n = 1 
// 输出: [1,2,3,4,5,6,7,8,9]
// 说明：用返回一个整数列表来代替打印 n 为正整数

package main 
import ("fmt";
	"strconv"
)



func PrintLargeNumber( n int )string {
	if n==0{
		return ""
	}else{
		bigest:=""
		num:=0
		result:=""
		for i:=0 ;i<n;i++{
			bigest=bigest+"9"
			num, _ = strconv.Atoi(bigest)
		}
		for i:=1;i<=num;i++{
			result=result+" "+strconv.Itoa(i)
		}
		return result
	}
	return ""
}
func main(){
	n:=3
	fmt.Println(PrintLargeNumber(n))
}



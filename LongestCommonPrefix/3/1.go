//编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
package main 
import (
	"fmt";
	"strings"
)


func LongestCommonPrefix(a []string) string {
	//var result string
	benchmark:= a[0]
	for len(benchmark)!=0{
		tag:=0
		for _,v:= range a{

			if strings.Index(v,benchmark)!=0{
				benchmark=benchmark[:len(benchmark)-1]
				tag=1
				break
			}


		}
		if tag==0{
			return benchmark
		}
		

	}
	return ""
	
}
func main()  {
	a:=[]string{"fler","flow","flight",}
	result:=LongestCommonPrefix(a)
	fmt.Println(result)	
}
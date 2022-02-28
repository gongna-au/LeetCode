//编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
package main 
import (
	"fmt";
	"strings"
)


func LongestCommonPrefix(a []string) string {
	//var result string
	var i=0
	benchmark:= a[0]
	for i=0;i<len(benchmark);i++ {
		tag:=1
		for _,v := range a{
			if !strings.HasPrefix(v,benchmark[0:i]){
				tag=0
			}
			if tag==0{
				break
			}
		}

		if tag==0{
			break
		}

	}
	return benchmark[0:i-1]
	

	
}
func main()  {
	a:=[]string{"fower","flow","flight",}
	result:=LongestCommonPrefix(a)
	fmt.Println(result)	
}
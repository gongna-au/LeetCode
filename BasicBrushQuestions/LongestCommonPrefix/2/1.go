//编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
package main 
import (
	"fmt";
	"strings"
)


func LongestCommonPrefix(a []string) string {
	//var result string
	benchmark:= a[0]
	var i=len(benchmark)-1
	for i=len(benchmark)-1;i>=0 ;i-- {
		tag:=1
		for _,v:=range a{
			if strings.Index(v,benchmark)!=0{
				tag=0
			}
		}
		if tag==1{
			return benchmark
		}else{
			benchmark=benchmark[0:i-1]
			continue
		}
	}

	return ""
	
}
func main()  {
	a:=[]string{"ower","flow","flight",}
	result:=LongestCommonPrefix(a)
	fmt.Println(result)	
}
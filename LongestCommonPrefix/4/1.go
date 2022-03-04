//编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
package main 
import (
	"fmt";
	"strings"
)


func LongestCommonPrefix(a []string) string {
	//var result string
	benchmark:= a[0]

	for _,v:= range a{
		for {
			if strings.Index(v,benchmark)!=0{

				if len(benchmark)==0{
					return ""
				}else{
					benchmark=benchmark[0:len(benchmark)-1]
				}
				
			}

		}
		
	}
	return benchmark
	
	
}
func main()  {
	a:=[]string{"fler","flow","flight",}
	result:=LongestCommonPrefix(a)
	fmt.Println(result)	
}
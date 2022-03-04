/* 旋转字符串
给定两个字符串, A 和 B。A 的旋转操作就是将 A 最左边的字符移动到最右边。
例如, 若 A = 'abcde'，在移动一次之后结果就是'bcdea' 。
如果在若干次旋转操作之后，A 能变成B，那么返回True。 
输入: A = 'abcde', B = 'cdeab'
输出: true

输入: A = 'abcde', B = 'abced'
输出: false
*/
package main
import (
	"fmt";
	"strings"
)

func  RotatingStrings(str1 string, str2 string )bool{
	if len(str1)!=len(str2) || str1=="" {
		return false
	}
	str1Array:=[]rune(str1)
	i:=0
	for ; !strings.Contains(str1,str2) && i<len(str1Array)-1; i++{
		str1=str1+string(str1Array[i])

	}
	if i==len(str1Array){
		return false
	}	
	return true
	
	
}


func main(){
	str1:="abcde"
	str2:="cdeab"
	fmt.Println(RotatingStrings(str1,str2))
}
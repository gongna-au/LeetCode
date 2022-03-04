/*题目：实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，
在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回 -1。*/
// 输入: haystack = "hello", needle = "ll"
// 输出: 2
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
package main
import ("fmt"
		//"strings"
)
func strStr(haystack string, needle string)int{
	if needle==""{
		return 0
	}else{
		haystackArray:=[]rune(haystack)
		needleArray :=[]rune (needle)
		
		for i:=len(needleArray)-1; i<len(haystackArray); {
			if i>=len(haystackArray){
				return -1
			}

			j:=len(needleArray)-1
			for  ;j>=0;{
				if haystackArray[i]==needleArray[j]{
					i--
					j--
				}else{
					break
				}
				
			}
			if j==-1{
				return i+1
			}else{
				i=i+len(needleArray)
			}
			
		}
		return -1
	}
}
func main(){
	haystack := "hello"
	needle :="ll"
	fmt.Println(strStr(haystack,needle))

}


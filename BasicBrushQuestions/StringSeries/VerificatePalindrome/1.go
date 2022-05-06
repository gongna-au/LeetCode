/*验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明： 本题中，我们将空字符串定义为有效的回文串。


示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
*/
/*  a:=[]rune("y")
	b:=[]rune("a")
	c:=[]rune("0") 48~57合法 65～90合法  97~122
	d:=[]rune("9")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	48~57合法 65～90合法  97~122
*/
package main
import "fmt"

func VerificatePalindrome(str string) bool{
	tag:=0
	if str==""{
		fmt.Println("An empty string is a palindrome")
		return true
	}else{
	
		strArray:=[]rune (str)
		i:=0
		j:=len(strArray)-1

		for ;i!=j;{
			illegalOfi:= (strArray[i]>=48 && strArray[i]<=57) || (strArray[i]>=65 && strArray[i]<=90) || (strArray[i]>=97 && strArray[i]<=122)
			illegalOfj:= (strArray[j]>=48 && strArray[j]<=57) || (strArray[j]>=65 && strArray[j]<=90) || (strArray[j]>=97 && strArray[j]<=122)
			if 	!illegalOfj{
				j--
				continue
			}
			if !illegalOfi {
				i++
				continue

			}
			if strArray[i]==strArray[j] || strArray[i]==strArray[j]+32 || strArray[i]==strArray[j]-32{
				i++
				j--
			}else{
				
				return false
			}
	
		}
		if i>=j{
			tag=1
		}

	}
	if tag==1{
		return true 

	}else{
		return false
	}
	
}

func main(){
	
	fmt.Println(VerificatePalindrome("race a car"))
	fmt.Println(VerificatePalindrome("A man, a plan, a canal: Panama"))

}


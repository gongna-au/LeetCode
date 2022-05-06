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

import "fmt"
func RotatingStrings( str1 string, str2 string )bool{
	str1Array:=[]rune(str1)
	str2Array:=[]rune(str2)

	str1Lenth:=len(str1Array)
	str2Lenth:=len(str2Array)
	if str1Lenth!=str2Lenth {
		return false
	}else{

		head1:=0
		tail1:=str1Lenth-1
		

			for {

				tail2:=str1Lenth-1
				for  tail2>=0{
					if str1Array[tail1]==str2Array[tail2]{
						tail2--
						tail1=(tail1+str1Lenth-1)%str1Lenth
					}else{
						break
					}
				}

				if tail2==-1{
					return true
				}else{

					head1=(head1+1)%str1Lenth
					tail1=(head1+(str1Lenth-1))%str1Lenth
				}

				if head1==0{
					break
				}
	
			}
			if head1 ==0{
				return false
			}


	}
	 return false



}
func main(){
	str1:="abcde"
	str2:="abced'"
	fmt.Println(RotatingStrings(str1,str2))

}
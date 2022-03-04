/*
最后一个单词的长度
给定一个仅包含大小写字母和空格 ' ' 的字符串 s，返回其最后一个单词的长度。如果字符串从左向右滚动显示，那么最后一个单词就是最后出现的单词。
示例：

输入: "Hello World" 
输出: 5
说明： 一个单词是指仅由字母组成、不包含任何空格字符的 最大子字符串。 
*/
package main
import ("fmt";
)



func LengthOfLast(str string) int {
	jump:=[]rune(" ")
	strArray:=[]rune(str)
	lastLenth:=0
	for i:=len(strArray)-1; ;{
		if strArray[i]==jump[0]{
			i--
		}else{

			for j:=i;strArray[j]!=jump[0];j--{
				lastLenth++
			}
			return lastLenth

		}

	}
	return -1
	
}
func  main()  {
	fmt.Println(LengthOfLast("Hello World"))
}



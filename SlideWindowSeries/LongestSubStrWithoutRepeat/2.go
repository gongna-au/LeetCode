package main

import (
	"fmt";
	"strings"

)
/*
无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

输入: "abcadc"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
*/

func LongestSubStrWithoutRepeat(str string)( result []string, lenth int ){
    strArray  := []byte(str)
	variableStr := "" 

	for i := 0; i < len(strArray) ;i++{
		temp := string( strArray[i] )
		if ( strings.Contains(variableStr,temp) == true ){
			index := strings.Index(variableStr,temp)
			variableStr = variableStr[index+1: len(variableStr)]
			variableStr=variableStr+temp
		}else{
			variableStr=variableStr+temp
		}
		if ( len(variableStr) >= lenth  ){
			if len(variableStr) > lenth {
				result = []string{}
				lenth = len(variableStr)
			}
			
			result = append(result,variableStr)
			
		}
		

	}

	return result,lenth
	
}
func main(){

	str := "abcadc"
	result,max := LongestSubStrWithoutRepeat(str)  
	fmt.Println(result)
	fmt.Println(max)

} 
// map的输出不是按照插入顺序输出的
// 所以在遍历删除的时候，删除的方式是错误的
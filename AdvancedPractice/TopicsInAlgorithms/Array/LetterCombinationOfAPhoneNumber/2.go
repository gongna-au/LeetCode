/*
	Leetcode 17.Letter Combinations of a Phone Number
	题⽬⼤意
	给定⼀个仅包含数字 2-9 的字符串，返回所有它能表示的字⺟组合。给出数字到字⺟的映射如下（与电话按键相同）。
	注意 1 不对应任何字⺟
	Input: "23"
	2-abc
	3-def
	4-ghi
	5-jkl
	6-mno
	7-pqrs
	8-tuv
	9-wxyz
	Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

*/
// 解法⼆ ⾮递归
package lettercombinationofaphonenumber

import (
//"fmt"
//"strings"
//"strconv"
)

func findCombine(digits string) []string {
	if len(digits) == 0 {
		return result
	}
	layers := mapList[int(digits[0]-'0')]
	temp := []string{}
	for i := 0; i < len(layers); i++ {
		if len(result) == 0 {
			result = append(result, "")
		}
		for j := 0; j < len(result); j++ {
			temp = append(temp, result[j]+string(layers[i]))
		}
	}
	result = temp
	findCombine(digits[1:])
	return result

	/* if len(digits) == 0 {
		return result
	}
	letter := mapList[int(digits[0]-'0')]
	temp := []string{}
	for i := 0; i < len(letter); i++ {

		if len(result) == 0 {
			result = append(result, " ")
		}

		for j := 0; j < len(result); j++ {
			temp = append(temp, result[j]+string(letter[i]))
		}

	}
	result = temp
	findCombination(digits[1:])
	return result */

}

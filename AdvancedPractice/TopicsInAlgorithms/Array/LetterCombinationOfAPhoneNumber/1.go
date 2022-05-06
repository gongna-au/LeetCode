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
// 解法⼆ DFS
package lettercombinationofaphonenumber

import (
//"fmt"
//"strconv"
)

var (
	mapList = map[int]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}
	result = []string{}
)

func LetterCombination(digits string) []string {
	if digits == "" {
		return []string{}
	}
	findCombination(&digits, 0, "")
	return result

}
func findCombination(digits *string, index int, trace string) {
	if index == len(*digits) {
		result = append(result, trace)
		return
	}

	letter := mapList[int((*digits)[index]-'0')]
	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, trace+string(letter[i]))
	}
	return

}

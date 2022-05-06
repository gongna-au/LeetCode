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
// 解法三 回溯
package lettercombinationofaphonenumber

import (
	"fmt"
	//"strings"
	//"strconv"
)

var (
	dict = map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	endResult = []string{}
)

func FindCombine(digits string) []string {
	if digits == "" {
		return endResult
	}
	layersBackTracking("", digits)
	return endResult

}
func layersBackTracking(temp string, digits string) {
	if len(digits) == 0 {
		endResult = append(endResult, temp)
		return
	}
	layersSlice := dict[string(digits[0])]
	digits = digits[1:]
	for _, v := range layersSlice {
		temp = temp + v
		layersBackTracking(temp, digits)
		temp = temp[0 : len(temp)-1]

	}

}
func main() {
	digits := "23"
	fmt.Println(FindCombine(digits))
}

/*


 */

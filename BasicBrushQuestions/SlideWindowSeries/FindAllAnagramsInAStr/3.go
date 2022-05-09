package main
import (
	"fmt";
	"strings"
)

/* 
找到字符串中所有字母异位词
给定一个字符串 s 和一个非空字符串 p，
找到 s 中所有是 p 的字母异位词的子串，
返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
说明：
字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。
示例 

输入:s: "cbaebabacd" p: "abc"
输出:[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。

输入:s: "abab" p: "ab
输出:[0, 1, 2]
解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词

*/
func FindAllAnagramsInAStr(str string ,pstr string) (resultIndex []int, result string) {
	strArray := []byte (str)
	pstrArray := []byte (pstr)
	//mapstrList := map[string]int {}
	//mappstrList := map[string]int {}
	

	
	for j := len(pstrArray)-1 ;j < len (strArray) ;j++{
		i:= j -(len(pstrArray)-1)
		temp :=""
		for  ;i <= j;i++{
			temp = temp+string(strArray[i])
			if (!strings.Contains(pstr,string(strArray[i]))) {
				break
			}
			
		}
		if i==j+1{
			resultIndex = append(resultIndex,j -(len(pstrArray)-1))
			result = result+" "+temp

		}

		
	}
	return resultIndex,result



}

func main(){
	str := "abab"
	pstr := "ab"
	fmt.Println(FindAllAnagramsInAStr(str,pstr))

}

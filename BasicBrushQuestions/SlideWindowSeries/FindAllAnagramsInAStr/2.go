package main
import (
	"fmt";
	//"strings"
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
	mappstrList := map[string]int {}
	

	for i := 0; i<len(pstrArray) ;i++ {
		mappstrList[string(pstrArray[i])]=mappstrList[string(pstrArray[i])]+1
	}
	

	for i := len(pstrArray)-1 ; i < len(strArray) ; i++{
		tempList := map[string]int {}

		for  j := i -(len(pstrArray)-1) ;j <= i ; j++  {
			tempList[string(strArray[j])] =  tempList[string(strArray[j])]+1
		}
		j := i -(len(pstrArray)-1)
		for  ;j <= i ; j++  {
			if mappstrList[string(strArray[j])] !=  tempList[string(strArray[j])] {
				break
			}
			
		}

		if j == i+1 {
			resultIndex = append(resultIndex,i-(len(pstrArray)-1))
			for k := i -(len(pstrArray)-1) ;k <= i ; k++  {
				result =  result+string(strArray[k])
			}
			result =  result+" "
		}


	}
	return resultIndex, result




}

func main(){
	str := "abab"
	pstr := "ab"
	fmt.Println(FindAllAnagramsInAStr(str,pstr))

}


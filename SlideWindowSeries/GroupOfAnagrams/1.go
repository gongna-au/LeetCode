package main
import (
	"fmt";
	//"strings"
)
/* 给你一个字符串数组，把字母异位词组合在一起，可以按照任意顺序返回结果列表 */
/* str = ["eat","tea","tan","ate","nat","bat"] */
func GroupOfAnagrams(strArray []string)(result []string){

	for i:=0;i<=len(strArray)-1;{
		resultMapList :=GetWord(strArray[i])
		str := strArray[i]
		strArray =strArray[i+1:]
		result = append(result,str)


		for j:=0;j<len(strArray)-1;{
			 if IfLeagal(strArray[j],resultMapList){
				result = append(result,strArray[j])
				strArray = append(strArray[:j],strArray[j+1:]...)
				 //fmt.Println(strArray)
				 continue

			 }else{
				j++
			 }

			
		}
		result = append(result,"/")

		


	}
	return result

}
func GetWord(str string) ( map[string] int){
	resultMapList := map[string]int { }
	array := []byte(str)
	for i:=0;i<len(array);i++{
		resultMapList[string(array[i])] = resultMapList[string(array[i])]+1
	}
	return resultMapList

}
func IfLeagal(str string,mapList map[string] int) bool{
	array := []byte (str)
	for i:=0;i<len(array);i++{
		_, ok :=mapList[string(array[i])]

		if !ok{
			return false
		}
	}
	return true
	


}

func main(){
	str := []string {"eat","tea","tan","ate","nat","bat"}
	fmt.Println(GroupOfAnagrams(str))



}


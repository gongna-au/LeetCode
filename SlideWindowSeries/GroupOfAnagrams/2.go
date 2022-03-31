package main
import (
	"fmt";
	//"strings"
)
/* 给你一个字符串数组，把字母异位词组合在一起，可以按照任意顺序返回结果列表 */
/* str = ["eat","tea","tan","ate","nat","bat"] */
func GroupOfAnagrams(strArray []string)(result []string){
	mapList := map[string]int{}
	k := 97
	for i:=1;i<=26;i++{
		
		mapList[string(k)]=i
		k++
		
	}
	

	
	for i:=0;i<=len(strArray)-1;{
		tag := Calculate(strArray[i],mapList)
		result = append(result,strArray[i])
		strArray =strArray[i+1:]
		for j:=0;j<len(strArray)-1;{
			if	Calculate(strArray[j],mapList) == tag {
				result = append(result,strArray[j])
				strArray = append(strArray[:j],strArray[j+1:]...)
				 
				 continue
			}else{
				j++
			}
		
		}
		result = append(result,"/")
		
		
	}
	return result

}
func Calculate(str string , mapList map[string]int) int{
	array :=[]byte (str)
	sum := 1
	for _,v := range array{
		sum = sum *mapList[string(v)]
	}
	return sum	
}


func main(){
	str := []string {"eat","tea","tan","ate","nat","bat"}
	
	fmt.Println(GroupOfAnagrams(str))



}


package main
import (
	"fmt";
	"regexp";
	"strings"
)


func VerificatePalindrome( str string) bool{
	r, _ := regexp.Compile(`[A-Za-z0-9]`)
	result:=r.FindAllString(str,-1)
	//fmt.Println(result)
	target :=""
	for _,v:=range result{
		target=target+v
	} 
	targetArray:=[]rune(target)
	lenth:=len(targetArray)-1
	i:=0
	for ;i<lenth;i++{
		left :=strings.ToLower(string(targetArray[i]))
		right:=strings.ToLower(string(targetArray[lenth-i]))
		 if left !=right {
			 break
			 
		 }else{
			 continue
		 }
	}
	if i==lenth{
		return true

	}else{
		return false
	}
	

}
func main(){
	fmt.Println(VerificatePalindrome("race a car"))
	fmt.Println(VerificatePalindrome("A man, a plan, a canal: Panama"))
	//fmt.Println(strings.ToLower("A"))


}


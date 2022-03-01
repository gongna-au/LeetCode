package main
import "fmt"

func BFSearch( target string , model string )int {
	targetLen:=len(target)
	modelLen:=len(model)

	for  i:=0;i<targetLen ;i++{
		
		k:=i
		j:=0
		for ;j< modelLen;j++{
			
			
				
				if target[k]==model[j]{
					k++
					j++
				}else{
					
					break
				}
			
		}
		
		if j==modelLen{
			return i 
		}
	}
	return -1

}
func main(){
	target:="ababdcd"
	model:="cdf"
	fmt.Println(BFSearch(target,model))
}
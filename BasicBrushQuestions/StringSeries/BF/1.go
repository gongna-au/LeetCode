package main
import "fmt"

func BFSearch( target string , model string )int {
	targetLen:=len(target)
	modelLen:=len(model)
	for i:=0;i<targetLen; {
		if target[i]==model[0]{
			
			j:=1
			for ;j<modelLen;j++{
				if (i+j)>=targetLen{
					fmt.Println("Could not find subscript for string matching pattern")
					return -1
				}
				if target[i+j]!=model[j]{
					break

				}
			}
			if j==modelLen{
				return i
			}else{
				i++
			}

		}else{
			i++
		}
	}
	fmt.Println("Could not find subscript for string matching pattern")
	return -1

}
func main(){
	target:="ababdcd"
	model:="dcd"
	fmt.Println(BFSearch(target,model))
}
package main
import( "fmt";
		"sort"
	)

func intersection(a []int, b []int )[]int {
	sort.Ints(a)
	sort.Ints(b)
	i:=0
	j:=0
	k:=0
	for i>= len(a) && j>= len(b){
			
		if a[i]<b[j] {
			i++
		}
		
		if a[i]>b[j] {
			j++		
		}	
		if a[i]==b[j]{
			a[k]=a[i]
			k++
			i++
		}
	}
	return a[0:k]
}
func main(){
	a:=[]int{2,2,2}
	b:=[]int{2,3,4,5}
	c:=intersection(a,b)
	fmt.Println(c)
}
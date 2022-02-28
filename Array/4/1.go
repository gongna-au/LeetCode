// 题目在进阶问题中问道：如果给定的数组已经排好序呢？你将如何优化你的算法？
// 我们分析一下，假如两个数组都是有序的，分别为：arr1 = [1,2,3,4,4,13]，arr2 = [1,2,3,9,10]
package main
import "fmt"
func intersection(a []int ,b []int) []int  {
	k:=0
	j:=0
	for _,v:=range a {
		if  (j+1)<len(b)&& b[j]==b[j+1]{
			j++
		}
		if v==b[j]{
			a[k]=v
			k++
		}
		if v!=b[j]{
			j++
			if  v==b[j]{
				a[k]=v
				k++
			}
		}
	}
	
	return a[0:k]

}
func main(){
	a:=[]int{2,2,2,3,5}
	b:=[]int{2,2,5}
	c:=intersection(a,b)
	fmt.Println(c)

}
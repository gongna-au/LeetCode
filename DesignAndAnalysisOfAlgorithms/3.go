/* 两个序列的中位数字 */
package main
import "fmt";
func MedianOfSequence(s1 []int ,s2 []int) int {
	mid := (len(s1)-1)/2
	
	if s1[mid]== s2[mid]  {
		return s1[mid]
	}else if s1[mid] < s2[mid] {
		s1=s1[mid:]
		s2=s2[:mid+1]
		return 	MedianOfSequence(s1,s2)


	}else {
		s2=s2[mid:]
		s1=s1[:mid+1]
		return 	MedianOfSequence(s1,s2)

	}
	
}
func main(){
	s1 := []int {1, 2, 3, 4, 5}
	/* 3 4 5
	4 5 6
	4 5
	4 5
	4 
	4 
	 */
	s2 := []int {4, 5, 6, 7, 8}
	fmt.Println(MedianOfSequence(s1,s2))


}


/* 
按奇偶排序数组
给定一个非负整数数组 A，
返回一个数组，
在该数组中， 
A 的所有偶数元素之后跟着所有奇数元素。
你可以返回满足此条件的任何数组作为答案。
示例：
1 输入：[3,1,2,4]
2 输出：[2,4,3,1]
3 输出 [4,2,3,1]，
[2,4,1,3] 和 [4,2,1,3] 也会被接受。
提示：
1 <= A.length <= 5000
0 <= A[i] <= 5000 */
package main
import "fmt"
func SortArrayByParity(array []int)[]int{
	for i:=0; i<len(array)-1 ;{
		j:=0
		if array[i]%2 == 0 {
			i++
			continue

		}else{
			j = i+1

		}
		tag :=0
		for ;j<len(array); {
			if array[j] %2 == 0 {
				tag = 1
				break
			}else{
				j++

			}

		}
		if tag == 1{
			array[i],array[j] = array[j],array[i]
		}

		i++

	}
	return array

}

func main(){
	array := []int {3,1,7,5,2,8,4}
	fmt.Println(SortArrayByParity(array))


}




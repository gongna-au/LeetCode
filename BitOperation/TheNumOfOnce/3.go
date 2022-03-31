/* 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只
出现了一次的元素。说明：你的算法应该具有线性时间复杂度。你可以不使用额外空间来实现吗？ */
package main

import "fmt"
func NumberAttendOnce(array []int )( result int){
	//去重 然后*3 = 该值的两倍
	//用空结构体节省空间
	j := 0
	mapList := map[int] struct{} {}
	sum1:=0
	for i:=0;i < len(array);i++{
		sum1 = sum1+array[i]
		_,ok := mapList[array[i]]
		if ok {
			continue

		}else{
			mapList[array[i]]=struct {}{}
			array[j] = array[i]
			j++
		}


	}
	//不要忘记这一步
	array=array[:j]

	sum2:=0
	for _,v := range array{
		sum2 = sum2 + v

	}
	
	return (3*sum2-sum1)/2

}
func main(){
	array1 := []int {2,2,1,2,3,3,3,4,4,4}
	array2 := []int {5,2,2,2,3,3,3}
	fmt.Println(NumberAttendOnce(array1))
	fmt.Println(NumberAttendOnce(array2))


}
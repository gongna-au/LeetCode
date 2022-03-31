package main
import "fmt"
/* 
阿珂喜欢吃香蕉
这里总共有 N 堆香蕉，第 i 堆中有piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。 阿珂可
以决定她吃香蕉的速度 K （单位：根/小时），每个小时，她将会选择一堆香蕉，从中吃掉 K 根。
如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。返回她可以在 H 小时内吃掉所有香蕉的最小
速度 K（K 为整数）。
示例 1：1 输入: piles = [3,6,7,11], H = 8
输出: 4
//9 0 0 0 1 0
//0 0 0 1 0
//0-3 1—0 1-1  1-5  
          0-1  0-5
// 0-3 1-1 1-2 2-1 
   1   1 0-1  1  0-2  2 0-1
// 0-3 1-2 1-3 2-3
  1+  1 + 0-2 + 1 + 0-3 + 2 + 0-3 
示例 2：
输入: piles = [30,11,23,4,20], H = 5
输出: 30
示例 3：
输入: piles = [30,11,23,4,20], H = 6
输出: 23 
*/
func KekeWhoLovesBananas(array []int ,H int )(result int ) {

	if len(array) > H {
		fmt.Println("Keke can not eat all the bananas")
		return -1
	}
	maxSpeed := array[0]
	for _,v := range array {
		if v > maxSpeed {
			maxSpeed = v
		}
	}
	remainder :=0
	quotient :=0
	lowSpeed := 1
	highSpeed := maxSpeed
	midSpeed := (highSpeed + lowSpeed)/2

	T :=0
	// 1 6 11 
	// 
	for  {

		T= 0
		for _,v := range array{
			quotient = v/midSpeed
			remainder= v%midSpeed
			
			if quotient == 0 {
				T= T+1
			}else if quotient != 0 && remainder != 0 {
				T= T+quotient+1
			}else if quotient != 0 && remainder == 0{
				T= T+quotient
			}
				
		}
		if T > H {
			//吃不完 吃得太慢
			//fmt.Println("吃不完 吃得太慢")
			lowSpeed =  midSpeed+1
			midSpeed = (highSpeed + lowSpeed)/2
			
		}else if T <= H{
			//吃的完
			//fmt.Println("吃的完")
			highSpeed = midSpeed
			midSpeed = (highSpeed + lowSpeed)/2
			//fmt.Println(midSpeed)
			
		}

		if lowSpeed == highSpeed{
			return lowSpeed
		}
		

	}
	
	
	
}
func main(){
	array1 := []int {3,6,7,11}
	array2 := []int {30,11,23,4,20}
	H1 := 8
	H2 := 6

	fmt.Println(KekeWhoLovesBananas(array1,H1))
	fmt.Println(KekeWhoLovesBananas(array2,H2))


}


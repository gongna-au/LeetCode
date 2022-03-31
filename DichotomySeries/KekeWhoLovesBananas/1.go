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
	minSpeed := array[0]
	for _,v := range array {
		if v < minSpeed {
			minSpeed = v
		}
	}
	//fmt.Println(minSpeed)
	
	tag := 1
	lastT := 0
	increase := false
	down := false

	for {

		T := 0
		for _,v := range array{
			quotient := v / minSpeed
			//fmt.Println(quotient)
			remainder := v % minSpeed
			//fmt.Println(remainder)
			if remainder != 0{
				T = T+quotient + 1
			}else{
				T = T+quotient 
			}
		}
		//上次吃不完，这次刚好吃的完
		if increase && T <= H{
			
			return minSpeed
		}
		//上一次吃的完，这一次吃不完
		if down && T >H{
			
			return lastT
		}

		if T > H {
			//吃不完 
			lastT = T
			increase = true
			//加速
			minSpeed = minSpeed+1
			tag = 1
		}else if T <= H{
			// 吃的完

			down = true
			lastT = T
			minSpeed = minSpeed-1
			tag =1
		}
		if tag == 1 {
			continue
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


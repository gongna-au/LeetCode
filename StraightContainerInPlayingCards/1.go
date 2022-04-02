package main

import (
	"fmt"
	"sort"
)

/*
扑克牌中的顺子
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A
为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。示例 1:
1 输入: [1,2,3,4,5]
2 输出: True
示例 2:
1 输入: [0,0,1,2,5]
2 输出: True
限制：
1 数组长度为 5
2 数组的数取值为 [0, 13]*/
func StraightContainerInPlayingcards(array []int) bool {
	sort.Ints(array)
	index := 0
	for k, v := range array {
		if v != 0 {
			index = k
			break

		}
	}
	distance := array[len(array)-1] - array[index]
	if distance == 4 {

		return true
	} else if distance < 4 {
		if distance > index-1 {

			return true
		} else {

			return false
		}
	} else {

		return false
	}
}

func main() {
	array := []int{0, 9, 11, 12, 13}
	fmt.Println(StraightContainerInPlayingcards(array))

}

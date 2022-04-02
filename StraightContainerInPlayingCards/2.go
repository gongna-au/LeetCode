package main

import (
	"fmt"
	//"runtime/trace"
	//"sort"
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
// 10 12 0 11 9
func StraightContainerInPlayingcards(array []int) bool {
	min, max := array[0], array[0]
	num0 := 0
	flag := 0
	for _, v := range array {

		if v == 0 {
			num0++
		}
		if v != 0 && flag == 0 {
			min, max = v, v
			flag = 1
		}
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	distance := max - min
	if distance == 4 {
		return true
	} else if distance > 4 {
		if distance-4 == num0 {
			return true
		} else {
			return false
		}
	} else {
		return true
	}

}

func main() {
	array := []int{0, 9, 11, 12, 13}
	fmt.Println(StraightContainerInPlayingcards(array))
	// 0 0 6 8 9
}

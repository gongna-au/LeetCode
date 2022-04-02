package main

import (
	"fmt"
	"sort"
)

/*
题目——1033
三枚石子放置在数轴上，位置分别为 a，b，c。每一回合，我们假设这三枚石子当前分别位于位置
x, y, z 且 x < y < z。从位置 x 或者是位置 z 拿起一枚石子，并将该石子移动到某一整数位置 k 处，
其中 x < k < z 且 k != y。当你无法进行任何移动时，即，这些石子的位置连续时，游戏结束。
要使游戏结束，你可以执行的最小和最大移动次数分别是多少？以长度为 2 的数组形式返回答案：
answer = [minimum_moves, maximum_moves]
示例1：
1 输入：a = 1, b = 2, c = 5
2 输出：[1, 2]
3 解释：将石子从 5 移动到 4 再移动到 3，或者我们可以直接将石子移动到 3。
// 1  5  7
1      6      9
示例 2：
1 输入：a = 4, b = 3, c = 2
2 输出：[0, 0]
3 解释：我们无法进行任何移动。
提示：
1 1 <= a <= 100
2 1 <= b <= 100
3 1 <= c <= 100
4 a != b, b != c, c != a */
//    3   6   10

func MoveStonesUntilContinuous(array []int) []int {
	sort.Ints(array)
	result := make([]int, 2)
	if array[2]-array[0] == 2 {
		result[0] = 0
		result[1] = 0
		fmt.Println("我们无法进行任何移动")
		return result
		// 4 5 4 7 10
	} else {
		distance1 := array[1] - array[0]
		distance2 := array[2] - array[1]
		if distance1 == 1 || distance2 == 1 {
			result[0] = 1
		}
		if distance1 >= 2 {
			//移连续
			//再并上
			result[0] = 1 + 1
		}

		result[1] = (array[2] - (array[1] + 1)) + ((array[1] - 1) - array[0])
		return result

	}

}
func swap(array []int, i int, j int) {

	array[i], array[j] = array[j], array[i]

}
func main() {
	array := []int{1, 2, 5}
	fmt.Println(MoveStonesUntilContinuous(array))

}

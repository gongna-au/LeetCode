package main

import (
	"fmt"
	"sort"
)

/*
输入：people = [1,2], limit = 3
输出：1
解释：1 艘船载 (1, 2)
输入：people = [3,5,3,4], limit = 5
输出：4
解释：4 艘船分别载 (3), (3), (4), (5)
3 3 2 4 6
*/
func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func Lifeboats(array []int, limit int) int {
	sort.Ints(array)
	i := 0
	j := len(array) - 1
	num := 0
	for i <= j {
		if i == j {
			num = num + 1
			fmt.Printf("(%d)", array[i])
			fmt.Println()
			break
		} else if (array[i] + array[j]) <= limit {

			fmt.Printf("(%d,%d)", array[i], array[j])
			num = num + 1
			i++
			j--
		} else {
			//体重最重的人一定是单独乘坐一艘船

			fmt.Printf("(%d)", array[j])
			j--
			num = num + 1
		}

	}
	return num

}
func deleteSpeficOneIndexInarray(array []int, index int) []int {
	array = append(array[0:index], array[index+1:]...)
	return array
}
func deleteSpeficTWOIndexInarray(array []int, index1 int, index2 int) []int {
	array = append(array[0:index1], array[index1+1:]...)
	array = append(array[0:index2-1], array[index2:]...)
	return array
}

func main() {
	array := []int{3, 5, 3, 4}
	// 3 3 4 5

	num := Lifeboats(array, 5)
	fmt.Println("need boat is =", num)

}

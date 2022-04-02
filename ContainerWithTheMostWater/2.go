package main

/*
盛最多水的容器
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直
线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容
器可以容纳最多的水。 */
import (
	"fmt"
	//"sort"
)

func Containerwiththemostwater(array []int) int {
	result := 0
	for i := len(array) - 1; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if array[j] < array[i] {
				continue
			} else if array[i]*(i-j) > result {
				result = array[i] * (i - j)

			}

		}

	}

	return result

}
func main() {
	array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(Containerwiththemostwater(array))

}

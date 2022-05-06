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
	i := 0
	j := len(array) - 1
	for {
		if j > i {
			if array[j] > array[i] {
				if (j-i)*array[i] > result {
					result = (j - i) * array[i]
				}

				i++
			} else {
				if (j-i)*array[i] > result {
					result = (j - i) * array[j]
				}
				j--

			}
		} else {
			break
		}

	}

	return result

}
func main() {
	array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(Containerwiththemostwater(array))

}

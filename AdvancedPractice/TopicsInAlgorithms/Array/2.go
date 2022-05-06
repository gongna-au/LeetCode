/* LeetCode-15
给定⼀个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合。
输入 [-1, 0, 1, 2, -1, -4],
输出
	[-1, 0, 1],
	[-1, -1, 2]
	//-2 0 -2 3
	//-4 -1 -1 0 1 2
	// 0  1  2 3 4 5

*/
package Array

import (
	"fmt"
	"sort"
	//"strconv"
)

// 解法2
func ThreeSum(array []int) (result [][]int) {
	sort.Ints(array)
	fmt.Println(array)
	for i := 1; i < len(array)-1; i++ {
		for array[i] == array[i+1] {
			i++
		}
		//fmt.Println("i", i)

		start := 0
		end := len(array) - 1
		for start < i && i < end {

			if array[start] == array[start+1] && start+1 < i {
				start++
				continue
			}
			if array[end-1] == array[end] && end-1 > i {
				end--
				continue
			}

			if (array[start] + array[i] + array[end]) == 0 {
				result = append(result, []int{array[start], array[i], array[end]})
				start++
				end--

			} else if (array[start] + array[i] + array[end]) < 0 {
				start++

			} else if (array[start] + array[i] + array[end]) > 0 {
				end--
			}

		}

	}
	return result

}
func main() {
	array := []int{-1, 0, 1, 2, -1, -4}

	temp := ThreeSum(array)

	fmt.Println(temp)

}

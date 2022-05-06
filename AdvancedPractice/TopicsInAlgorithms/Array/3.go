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

// 解法3
func ThreeNumAdd(array []int) (result [][]int) {
	sort.Ints(array)
	fmt.Println(array)
	i := 0
	j := len(array) - 1
	uniqueArray := []int{}
	mapList := map[int]int{}
	//去掉数组中重复的元素
	for i < j {
		if array[i] == array[i+1] {
			mapList[array[i]] = mapList[array[i]] + 1
			i++
			continue
		}
		if array[j] == array[j-1] {
			mapList[array[j]] = mapList[array[j]] + 1
			j--
			continue
		}

		uniqueArray = append(uniqueArray, array[i], array[j])
		mapList[array[i]] = mapList[array[i]] + 1
		mapList[array[j]] = mapList[array[j]] + 1
		i++
		j--

	}
	uniqueArray = append(uniqueArray, array[i])
	mapList[array[i]] = mapList[array[i]] + 1

	//排序不能忘
	sort.Ints(uniqueArray)

	//利用uniqueArraay以及mapList来
	for i := 0; i < len(uniqueArray); i++ {
		if uniqueArray[i] == 0 && mapList[uniqueArray[i]] >= 3 {
			result = append(result, []int{uniqueArray[i], uniqueArray[i], uniqueArray[i]})

		}

		for j := i + 1; j < len(uniqueArray); j++ {
			if uniqueArray[i]*2+uniqueArray[j] == 0 && mapList[uniqueArray[i]] >= 2 {
				result = append(result, []int{uniqueArray[i], uniqueArray[i], uniqueArray[j]})
			}
			if uniqueArray[j]*2+uniqueArray[i] == 0 && mapList[uniqueArray[j]] >= 2 {
				result = append(result, []int{uniqueArray[i], uniqueArray[i], uniqueArray[j]})
			}
			c := -(uniqueArray[j] + uniqueArray[i])
			//保证下标不同
			_, ok := mapList[c]
			//c > uniqueArray[j] 就是保证你的这个数在后面
			if c > uniqueArray[j] && ok {
				result = append(result, []int{uniqueArray[i], uniqueArray[i], uniqueArray[j]})
			}

		}

	}

	for k, v := range mapList {
		fmt.Println("k", k)
		fmt.Println("v", v)
	}

	return result

}

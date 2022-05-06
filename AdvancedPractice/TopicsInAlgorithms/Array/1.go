package Array

/* LeetCode-15
给定⼀个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合。
输入 [-1, 0, 1, 2, -1, -4],
输出
	[-1, 0, 1],
	[-1, -1, 2]
	//-4 -1 -1 0 1  2
	//-4 -1 -1 0 1 2
	// 0  1  2 3 4 5
	// 解法1

*/

import (
	//"fmt"
	"reflect"
	"sort"
)

func SumOfThreeNumber(array []int) (result [][]int) {

	sort.Ints(array)
	AddList := map[int]int{}
	ElemList := map[int]int{}
	SameElemList := map[int]int{}
	for i := 0; i < len(array)-1; i++ {

		for j := i + 1; j < len(array); j++ {

			AddList[array[i]+array[j]] = AddList[array[i]+array[j]] + 1
		}
	}
	for i := 0; i < len(array); i++ {
		if ElemList[array[i]] == 0 {
			ElemList[array[i]] = i
		} else {
			SameElemList[array[i]] = i
		}
	}

	for i := 0; i < len(array); i++ {
		if _, ok := AddList[-array[i]]; ok {
			for j := i + 1; j < len(array); j++ {
				if k, ok := ElemList[-(array[i] + array[j])]; ok {
					if k != i && k != j {
						temp := []int{array[i], array[j], -(array[i] + array[j])}
						sort.Ints(temp)
						result = append(result, temp)
					} else {
						if k, ok = SameElemList[-(array[i] + array[j])]; ok {
							if k != i && k != j {
								temp := []int{array[i], array[j], -(array[i] + array[j])}
								sort.Ints(temp)
								result = append(result, temp)
							}

						}

					}
				}
			}
		}
	}

	return result

}
func DeleteSameElem(array [][]int) [][]int {
	for i := 0; i < len(array); i++ {
		sort.Ints(array[i])
	}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); {
			tag := reflect.DeepEqual(array[i], array[j])
			//fmt.Print(tag)
			if tag == true {

				array = append(array[0:j], array[j+1:]...)
				//fmt.Println(array)

			} else {
				j++
			}
		}

	}
	return array

}

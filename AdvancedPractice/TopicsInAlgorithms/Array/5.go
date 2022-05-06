/* 题⽬⼤意
给定⼀个数组，要求在这个数组中找出 3 个数之和离 target 最近。
*/
/*
Given array nums = [-1, 2, 1, -4],
and target = 1.
The sum that is closest to the target is 2. (-1 + 2 + 1 = 2)
// -3 -2         3
*/
package Array

import (
	//"archive/tar"
	//"fmt"
	"sort"
)

func ThreeSumClose(array []int, target int) (result [][]int) {
	// -4 -1 1 2
	// -5 -2 1 2 3
	// 距离越近距离
	/* c-(a) > c-(b) */
	//选择移动i会使得距离目标更近！！也就是选择后半部分

	// 这种情况只适用于重复元素数组不能全都是重复的元素
	//   2-3 2-4 2-5 2-6
	//   1-2 1-3 1-4 1-5

	sort.Ints(array)
	i := 0

	minNum := array[i] + array[i+1] + array[len(array)-1]
	for ; i < len(array)-2; i++ {
		k := len(array) - 1
		for j := i + 1; j < k; j++ {

			currentNum := array[i] + array[j] + array[k]
			if abs(target-currentNum) < abs(minNum-currentNum) {
				result = append([][]int{}, []int{array[i], array[j], array[k]})
			} else if abs(target-currentNum) == abs(minNum-currentNum) {
				result = append(result, []int{array[i], array[j], array[k]})
			}

		}

	}

	return result

}

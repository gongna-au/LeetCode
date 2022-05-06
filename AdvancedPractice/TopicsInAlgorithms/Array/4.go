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
	//"fmt"
	"sort"
)

func ThreeSumClosest(array []int, target int) (result [][]int) {
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
	j := len(array) - 1
	min := array[i] + array[j] + array[i+1]
	k := i + 1

	for i < j-1 {

		closest := array[i] + array[j] + array[k]

		if abs(target-closest) < abs(target-min) {
			result = append([][]int{}, []int{array[i], array[k], array[j]})
			min = array[i] + array[j] + array[k]

		} else if abs(target-closest) == abs(target-min) {
			result = append(result, []int{array[i], array[k], array[j]})
		}

		if k == j-1 {

			if abs(array[i+1]+array[j]+array[i+2]-target) < abs(array[i]+array[j-1]+array[i+1]-target) {
				i++
				k = i + 1
				continue
			} else if abs(array[i+1]+array[j]+array[i+2]-target) > abs(array[i]+array[j-1]+array[i+1]-target) {
				j--
				k = i + 1
				continue
			} else {
				//根据数学分析，我们知道了后半部分的和更加大，距离目标更近
				i++
				k = i + 1
				continue
			}
		}
		k++
	}
	return result

}
func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

/* 题目示例
第278题：第一个错误的版本
假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。 */
/* 你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。实现
一个函数来查找第一个错误的版本。你应该尽量减少对调用 API 的次数。示例:

给定 n = 5，并且 version = 4 是第一个错误的版本。

 调用 isBadVersion(3) -> false
 调用 isBadVersion(5) -> true
 调用 isBadVersion(4) -> true


所以，4 是第一个错误的版本。
这个题目还是相当简单的....我拿出来讲的原因，是因为我的开发生涯中，真的遇到过这样一件
事。当时我们做一套算薪系统，算薪系统主要复杂在业务上，尤其是销售的薪资，设计到数百个
变量，并且还需要考虑异动（比如说销售A是团队经理，但是下调到B团队成为一名普通销售，然
后就需要根据A异动的时间，来切分他的业绩组成。同时，最恶心的是，普通销售会影响到其团队
经理的薪资，团队经理又会影响到营业部经理的薪资，一直到最上层，影响到整个大区经理的薪
资构成）要知道，当时我司的销售有近万名，每个月异动的人就有好几千，这是非常非常复杂
的。然后我们遇到的问题，就是同一个月，有几十个团队找上来，说当月薪资计算不正确（放在
个人来讲，有时候差个几十块，别人也是会来找的）最后，在一阵漫无目的的排查之后，我们采
用二分的思想，通过切变量，最终切到错误的异动逻辑上，进行了修正。
*/
package main

import "fmt"

func isBadVersion(i int, array []bool) bool {
	if array[i] == false {
		return false
	}
	return true

}
func firstWrongVersion(array []bool) int {

	left := 0
	right := len(array) - 1

	for {
		mid := (left + right) / 2
		if isBadVersion(mid, array) == false {
			right = mid - 1

		} else {
			left = mid

		}
		if left <= right {
			break
		}

	}
	return left

}
func main() {
	array := []bool{true, true, true, true, true, true, false, false, false, false, false}
	fmt.Println(firstWrongVersion(array))

}

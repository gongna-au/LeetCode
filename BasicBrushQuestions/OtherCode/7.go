/*
灯泡开关（319）
题目示例
初始时有 n 个灯泡关闭。第 1 轮，你打开所有的灯泡。
第 2 轮，每两个灯泡关闭一次。
第 3 轮，
每三个灯泡切换一次开关（如果关闭则开启，如果开启则关闭）。
第 i 轮，每 i 个灯泡切换一次开关。
对于第 n 轮，你只切换最后一个灯泡的开关。
找出 n 轮后有多少个亮着的灯泡。
示例:
输入: 3
输出: 1
解释:
初始时, 灯泡状态 [关闭, 关闭, 关闭].
第一轮后, 灯泡状态 [开启, 开启, 开启].
第二轮后, 灯泡状态 [开启, 关闭, 开启].
第三轮后, 灯泡状态 [开启, 关闭, 关闭].

你应该返回 1，因为只有一个灯泡还亮着。 */
package main

import "fmt"

func LightBulbSwitch(n int) (lightNum int) {
	flag := make([]int, n)

	for k := range flag {
		flag[k] = 1
	}
	fmt.Println("i= 1", flag)

	for i := 2; i <= n; i++ {

		lengh := i - 1
		for lengh <= len(flag)-1 {

			//取反
			if flag[lengh] == 1 {
				flag[lengh] = 0
			} else {
				flag[lengh] = 1
			}
			fmt.Println("i=", i, flag)
			lengh = lengh + i
		}

	}
	lightNum = 0
	fmt.Println(flag)
	for k, _ := range flag {
		if flag[k] == 1 {
			lightNum++
		}
	}
	return lightNum

}
func main() {
	fmt.Println(LightBulbSwitch(12))

}

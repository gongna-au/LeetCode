package main

import "fmt"

/* 第299题：
猜数字（Bulls and Cows）游戏
你写下一个数字让你的朋友猜。每次他猜测后，你给他一个提示，告诉他有多少位数字和确切位置
都猜对了（称为“Bulls”, 公牛），有多少位数字猜对了但是位置不对（称为“Cows”, 奶牛）。你的朋
友将会根据提示继续猜，直到猜出秘密数字。
请写出一个根据秘密数字和朋友的猜测数返回提示的函数，用 A 表示公牛，用 B 表示奶牛。
请注意秘密数字和朋友的猜测数都可能含有重复数字
*/
func GuessTheNumber(array1 []int, array2 []int) {
	if len(array1) != len(array2) {
		fmt.Println("Input illegal!")
	}
	Bulls := 0
	Cows := 0
	secretMap := map[int]int{}
	guessMap := map[int]int{}
	for i := 0; i < len(array1); i++ {
		if array1[i] == array2[i] {
			Bulls++

		} else {
			secretMap[array1[i]] = secretMap[array1[i]] + 1
			guessMap[array2[i]] = guessMap[array2[i]] + 1

		}

	}
	for k, _ := range guessMap {
		if _, ok := secretMap[k]; ok {
			Cows++

		}
	}
	fmt.Println(Bulls, "A", Cows, "B")
}
func main() {
	array1 := []int{1, 1, 2, 3}
	array2 := []int{0, 1, 1, 1}
	GuessTheNumber(array1, array2)

}

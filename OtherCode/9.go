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
	j := 0
	for i := 0; i < len(array2); {
		if array2[i] == array1[j] {
			Bulls++
			j++
			array2 = append(array2[0:i], array2[i+1:]...)

		} else {
			j++
			i++
		}

	}
	//fmt.Println(array2)
	arrayList := map[int]int{}
	for k := 0; k < len(array1); k++ {
		arrayList[array1[k]] = arrayList[array1[k]] + 1
	}
	//fmt.Println(arrayList)
	//去重
	tempList := map[int]int{}
	for k := 0; k < len(array2); k++ {
		tempList[array2[k]] = tempList[array2[k]] + 1
	}
	fmt.Println(tempList)

	temp := []int{}
	for k, _ := range tempList {
		temp = append(temp, k)

	}
	fmt.Println(temp)

	for _, v := range temp {
		if _, ok := arrayList[v]; ok {
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

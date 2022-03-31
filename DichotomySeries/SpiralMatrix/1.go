package main

import "fmt"

/* 输入:

  [ 1, 2, 3 ],
  [ 4, 5, 6 ],
  [ 7, 8, 9 ]

7 输出: [1,2,3,6,9,8,7,4,5] */
func spiralmatrix(array [][]int) {
	row := len(array)
	list := len(array[0])
	index := [1][2]int{{0, 0}}
	//设置边界
	up := 0
	down := row - 1
	left := 0
	right := list - 1
	//开始移动

	for {

		tag := 1

		if tag == 1 {
			for index[0][1] <= right {
				fmt.Print(array[index[0][0]][index[0][1]])
				index[0][1]++
			}
			index[0][1]--
			index[0][0]++
			tag++

		}

		if tag == 2 {
			for index[0][0] <= down {
				fmt.Print(array[index[0][0]][index[0][1]])
				index[0][0]++
			}
			index[0][0]--
			index[0][1]--
			tag++

		}

		if tag == 3 {
			for index[0][1] >= left {
				fmt.Print(array[index[0][0]][index[0][1]])
				index[0][1]--
			}
			index[0][1]++
			index[0][0]--
			tag++

		}

		if tag == 4 {
			for index[0][0] >= up {
				fmt.Print(array[index[0][0]][index[0][1]])
				index[0][0]--
			}
			index[0][0]++

		}

		tag = tag % 4

		if index[0][0] == len(array)/2 {
			break
		}

		index[0][0]++
		index[0][1]++

		//边界缩小准备下一次循环
		up++
		down--
		left++
		right--

	}

}
func main() {

	array := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	spiralmatrix(array)

}

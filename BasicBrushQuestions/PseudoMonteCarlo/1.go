/* Pseudo Monte Carlo */
package main

import (
	//"crypto/x509"
	"fmt"
	//"math"
	"math/rand"
	"sort"
)

/* "crypto/rand"
"fmt" */

/* 算法介绍
期望：在概率论和统计学中，数学期望(mean)（或均值，亦简称期望）是试验中每次可能结果的
概率乘以其结果的总和，是最基本的数学特征之一。它反映随机变量平均取值的大小。
题目：在1*1的正方形中随机撒三个点，两两点都可构成长方形的一组对顶点，这样一共有三个长
方形，需要求面积第二大的长方形的面积的期望。
算法：每次随机三个点，计算第二大面积，最后统计期望。
02、蒙特卡洛
蒙特卡罗法也称统计模拟法、统计试验法。是把概率现象作为研究对象的数值模拟方法。是按抽
样调查法求取统计值来推定未知特性量的计算方法。蒙特卡罗是摩纳哥的著名赌城，该法为表明
其随机抽样的本质而命名。故适用于对离散系统进行计算仿真试验。在计算仿真中，通过构造一
个和系统性能相近似的概率模型，并在数字计算机上进行随机试验， */
func PseudoMonteCarlo(array [][]int) float64 {

	result := []int{}
	pointIndex := make([][]int, 3)
	for i := 0; i < 3; i++ {
		pointIndex[i] = make([]int, 2)
	}
	fmt.Println("Point", pointIndex)

	pointx := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			if array[i][j] == 1 {

				pointIndex[pointx][0] = i
				pointIndex[pointx][1] = j
				pointx = pointx + 1

			}

		}
	}

	x := Abs(pointIndex[0][0], pointIndex[1][0])
	y := Abs(pointIndex[0][1], pointIndex[1][1])
	result = append(result, x*y)
	x = Abs(pointIndex[1][0], pointIndex[2][0])
	y = Abs(pointIndex[1][1], pointIndex[2][1])
	result = append(result, x*y)
	x = Abs(pointIndex[0][0], pointIndex[2][0])
	y = Abs(pointIndex[0][1], pointIndex[2][1])
	result = append(result, x*y)
	sort.Ints(result)
	return float64(result[1]) / 9.0
}

func Abs(a int, b int) int {
	if a-b > 0 {
		return a - b
	} else {
		return b - a
	}

}
func RandomInitializate(array [][]int) {
	Xmaplistist := map[int]int{}
	Ymaplistist := map[int]int{}
	pointNum := 0
	for {
		x := rand.Intn(len(array))
		fmt.Println("X:", x)
		y := rand.Intn(len(array[0]))
		fmt.Println("Y:", y)

		_, ok1 := Xmaplistist[x]
		_, ok2 := Ymaplistist[y]
		if !ok1 && !ok2 {
			Xmaplistist[x] = 1
			Ymaplistist[y] = 1
			array[x][y] = 1
			pointNum++

		}
		if pointNum == 3 {
			break
		}

	}

}
func main() {
	array := [][]int{{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}}
	RandomInitializate(array)
	fmt.Println(array)
	fmt.Println(PseudoMonteCarlo(array))

}

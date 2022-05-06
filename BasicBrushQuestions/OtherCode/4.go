package main

import "fmt"

func productionOfWater(str string) bool {
	if len(str)%3 != 0 {
		//fmt.Print("1")
		return false
	}
	if GetNum(str, "H")/GetNum(str, "O") != 2 {

		return false
	}
	array := []byte(str)
	Harray := []byte("H")
	Oarray := []byte("O")
	flag := make([]bool, len(array))
	for k, _ := range flag {
		if array[k] == Harray[0] {
			flag[k] = true

		}
	}

	i1 := 0
	i2 := 0
	j := 0

	for j < len(flag) {

		if array[j] == Oarray[0] {
			//找到了O

			for i1 < len(flag) {

				if array[i1] == Harray[0] && flag[i1] == true {
					flag[i1] = false
					break
				} else {
					i1++
				}
				if i1 >= len(flag) {
					//fmt.Print("3")
					return false
				}

			}

			for i2 < len(flag) {
				if array[i2] == Harray[0] && flag[i2] == true {
					flag[i2] = false
					break
				} else {

					i2++
					if i2 >= len(flag) {
						//fmt.Print("4")
						return false
					}
				}

			}

			j++

		} else {
			j++
		}

	}
	return true

}
func GetNum(str string, target string) int {
	array := []byte(str)
	targetArray := []byte(target)
	num := 0
	for i := 0; i < len(array); i++ {
		if array[i] == byte(targetArray[0]) {
			num++

		}
	}
	return num

}
func main() {
	fmt.Println(productionOfWater("OHHOOO"))

}

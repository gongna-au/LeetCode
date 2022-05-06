package main

import "fmt"

type Node struct {
	value int
	Left  *Node
	Right *Node
}

func AdjustToBigRootHeap(array []*Node) []*Node {
	lastNodeIndex := len(array) - 1
	lastFatherNode := lastNodeIndex / 2
	for lastFatherNode > 0 {
		lastFatherNode = lastFatherNode - 1
		AdjustBigHeap(array, lastFatherNode)

	}
	return array

}

func AdjustBigHeap(array []*Node, fatherIndex int) {
	leftIndex := 2*fatherIndex + 1
	rightIndex := 2*fatherIndex + 2
	if fatherIndex >= len(array) || leftIndex >= len(array) {
		return
	}

	if array[leftIndex].value > array[fatherIndex].value {
		if array[leftIndex].value > array[rightIndex].value {
			array[fatherIndex].value, array[leftIndex].value = array[leftIndex].value, array[fatherIndex].value
			AdjustBigHeap(array, leftIndex)
		} else {
			array[fatherIndex].value, array[rightIndex].value = array[rightIndex].value, array[fatherIndex].value
			AdjustBigHeap(array, rightIndex)

		}
	} else if array[leftIndex].value < array[fatherIndex].value {
		if array[fatherIndex].value < array[rightIndex].value {
			array[fatherIndex].value, array[rightIndex].value = array[rightIndex].value, array[fatherIndex].value
			AdjustBigHeap(array, rightIndex)
		}
	}

}

func swapTwoNodeInArray(array []*Node, first int, second int) {
	if array[first].value < array[second].value {
		array[first].value, array[second].value = array[second].value, array[first].value
	}

}
func Adjust(father *Node) *Node {
	if father == nil {
		return nil
	}
	if father.Left == nil && father.Right == nil {
		return father
	}
	if father.Left != nil && father.Right != nil {

		if father.Right.value > father.value {
			if father.Right.value > father.Left.value {

				temp := father.Right
				father.Right = father
				father = temp
				Adjust(father.Right)
			} else {
				return father.Left

			}

		} else if father.Left.value > father.value {

			if father.Left.value > father.Right.value {

				return father.Left
			} else {
				return father.Right
			}

		}

	}
	return father

}

func swap(father *Node, child *Node) {
	father, child = child, father
	fmt.Println(father.value)
	fmt.Println(child.value)

}
func MaxChildNode(father *Node) (*Node, bool) {
	fmt.Println(3333)
	if father.Left != nil && father.Right != nil {
		fmt.Println(44444)

		if father.Right.value > father.value {
			if father.Right.value > father.Left.value {
				fmt.Println(99999)
				return father.Right, true
			} else {
				return father.Left, true

			}

		} else if father.Left.value > father.value {

			if father.Left.value > father.Right.value {

				return father.Left, true
			} else {
				return father.Right, true
			}

		} else {
			return father, true
		}

	} else if ((father.Left != nil) && (father.Right == nil)) && (father.Left.value > father.value) {

		return father.Left, true
	} else if ((father.Left == nil) && (father.Right != nil)) && (father.Right.value > father.value) {
		return father.Right, true
	} else {
		return father, false
	}

}
func main() {
	left1_left := &Node{
		3,
		nil,
		nil,
	}
	left1_Right := &Node{
		4,
		nil,
		nil,
	}

	left1 := &Node{
		1,
		nil,
		nil,
	}
	right1 := &Node{
		2,
		nil,
		nil,
	}

	root := &Node{
		0,
		nil,
		nil,
	}

	/*       0
	         /\
			1  2
			/\
		   3  4
	*/
	array := []*Node{root, left1, right1, left1_left, left1_Right}

	array = AdjustToBigRootHeap(array)

	for _, v := range array {

		fmt.Println(v.value)
	}

}

package main

import "fmt"

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {
	return &Node{data, nil, nil}
}

//用来创建大根堆
func CreateTreeByPreOrder(dataArray []int) (root *Node) {
	nodeArray := []*Node{}

	for _, v := range dataArray {

		nodeArray = append(nodeArray, NewNode(v))
	}

	root = nodeArray[0]
	for i := 0; i < len(dataArray); i++ {
		if 2*i+1 < len(dataArray) && 2*i+2 < len(dataArray) {
			nodeArray[i].Left = nodeArray[2*i+1]
			nodeArray[i].Right = nodeArray[2*i+2]

		} else if 2*i+1 < len(dataArray) && 2*i+2 >= len(dataArray) {
			nodeArray[i].Left = nodeArray[2*i+1]
			nodeArray[i].Right = nil

		} else if 2*i+1 >= len(dataArray) {
			nodeArray[i].Left = nil
			nodeArray[i].Right = nil

		}

	}
	return root

}

//先序遍历
func PreOrder(t *Node) {
	if t != nil {
		fmt.Printf("%v ", t.Data)
		PreOrder(t.Left)
		PreOrder(t.Right)

	}

}

//中序遍历
func InOrder(t *Node) {
	if t != nil {
		InOrder(t.Left)
		fmt.Printf("%v ", t.Data)
		InOrder(t.Right)
	}

}

//后序遍历
func PostOrder(t *Node) {
	if t != nil {
		PostOrder(t.Left)
		PostOrder(t.Right)
		fmt.Printf("%v ", t.Data)
	}

}

//层序遍历 用来验证建立的二叉树
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var queue = []*Node{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Data.(int))
			queue = queue[1:]
		}
		level++
	}
	return res
}
func CreateTreeArray(dataArray []int) (nodeArray []*Node) {
	nodeArray = []*Node{}

	for _, v := range dataArray {

		nodeArray = append(nodeArray, NewNode(v))
	}
	return nodeArray
}

func GetDeleteNodeFather(root *Node, target int, result *Node) *Node {

	if root != nil {
		if root.Left != nil && root.Left.Data.(int) == target {
			return root
		}
		if root.Right != nil && root.Right.Data.(int) == target {
			return root
		}

		result = GetDeleteNodeFather(root.Left, target, result)
		result = GetDeleteNodeFather(root.Right, target, result)

	}

	return result

}

func maxNode(left *Node, right *Node) *Node {
	if left == nil && right != nil {
		return right
	} else if left != nil && right != nil {
		return max(left, right)
	} else if left != nil && right == nil {
		return left
	}
	return nil
}

func max(left *Node, right *Node) *Node {
	if left.Data.(int) > right.Data.(int) {
		return left
	} else {
		return right
	}

}
func DeleteTragetNodeAndAdjust(fatherNode *Node, target int) (result *Node) {
	right := false
	left := false
	if fatherNode.Left != nil && fatherNode.Left.Data.(int) == target {
		left = true
	}
	if fatherNode.Right != nil && fatherNode.Right.Data.(int) == target {
		right = true
	}
	if right == true {
		//把右节点删除后调整为大根堆
		fatherNode.Right = AdjustToBigRootHeap(fatherNode.Right)

	}
	if left == true {
		//把左边节点删除后调整为大根堆
		fatherNode.Left = AdjustToBigRootHeap(fatherNode.Left)

	}
	return fatherNode

}
func AdjustToBigRootHeap(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {

		return nil
	}
	if root.Left == nil && root.Right != nil {
		return root.Right
	}
	if root.Right == nil && root.Left != nil {
		return root.Left
	}

	//删除根节点并调整为大根堆
	if root.Left != nil && root.Right != nil {
		if root.Left.Data.(int) >= root.Right.Data.(int) {
			root.Data = root.Left.Data.(int)
			root.Left = AdjustToBigRootHeap(root.Left)
		} else {
			root.Data = root.Left.Data.(int)
			root.Right = AdjustToBigRootHeap(root.Right)
		}

	}
	return root

}

func main() {
	dataArray := []int{100, 43, 26, 9, 33, 11, 21, 2}
	root := CreateTreeByPreOrder(dataArray)
	fmt.Println(levelOrder(root))
	result := root
	fatherNode := GetDeleteNodeFather(root, 9, result)
	/*
		测试函数结果
		if fatherNode != nil {
			fmt.Println(fatherNode.Data.(int))

		}
		fmt.Println(fatherNode.Right) */
	DeleteTragetNodeAndAdjust(fatherNode, 9)
	fmt.Println(levelOrder(root))

	/*
		测试函数结果
		temp := AdjustToBigRootHeap(fatherNode.Right)
		fmt.Println(temp)
		fmt.Println(temp.Left)
		fmt.Println(temp.Right) */

	/*     54                                                     54
		   /\                                                    /  \
		 32   12          ------->删除23这个节点之后  得到         32   12
		/\    /\                                               /\   /  \
	   8  23  5 9                                             8  3  5   9
	  /\ /\                                                  /\  / \
	 2 0 3 1                                                2 0 nil 1


	*/

	/*     100                                                    100
		   /\                                                    /  \
		 43   26          ------->删除9这个节点之后  得到         43   26
		/\    /\                                               / \   /  \
	   9  33  11 21                                          2  33  11  21
	  /
	 2



	*/

}

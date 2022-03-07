// 迭代解法

/* 二叉搜索树中的搜索
给定二叉搜索树（BST）的根节点和一个值。你需要在 BST 中找到节点值等于给定值的节点。返回以该节点为根的子树。如果节点不存在，则返回 NULL 。
示例：

给定二叉搜索树:

        4
       / \
      2   7
     / \
    1   3
    
和值: 2
你应该返回如下子树:

      2     
     / \   
    1   3 */
package main
import (
	"fmt";
	//"container/list"
)
type Tree struct{
	Data interface{}
	Left *Tree
	Right *Tree
}


type Initer interface{
	SetData( data interface{})
}

type Operater interface {
	PrintTree()
	PrintDepth() int 
	LeafCount() int 

}

type Order interface{
	PreOrder()
	InOrder()
	PostOrder()
}


func (t *Tree) SetData(data interface{}){

	t.Data=data

}
func (t *Tree) PrintTree(){

	PrintTree(t)	
}
func (t *Tree) PrintDepth()int {

	return  PrintDepth(t)
}

func (t *Tree) LeafCount() int {
	return LeafCount(t)
}

func (t *Tree) PreOrder() {
    PreOrder(t)
}

func (t *Tree) InOrder(){
	InOrder(t)
}

func (t *Tree) PostOrder(){
	PostOrder(t)
}

func PrintTree( t *Tree){
	if t != nil{
		fmt.Printf("%v",t.Data)
	}

	

	if t.Left != nil{
		fmt.Printf("( ")
		PrintTree(t.Left)
	}

	if t.Right != nil{
		fmt.Printf(",")
		PrintTree(t.Right)
		fmt.Printf(")")

	}

}

func max( a int , b int) int{
	if a >= b{
		return a

	}else{
		return b
	}
	
}

func PrintDepth(t *Tree ) int{
	var leftDepth int
	var rightDepth int

	if t == nil{
		return 0
	}else{
		leftDepth = PrintDepth(t.Left)
		rightDepth = PrintDepth(t.Right)

		if leftDepth>=rightDepth{
			return  leftDepth+1
		}else{
			return  rightDepth+1
		}
	}
}





func LeafCount(t *Tree )int{
	var leftLeaf int
	var rightLeaf int

	if t == nil{
		return 0
	}else if (t.Left == nil) &&  (t.Right == nil) {

		return 1

	}else  {
		leftLeaf = LeafCount(t.Left)
		rightLeaf = LeafCount(t.Right)

		

	}
	return  leftLeaf+rightLeaf
	

}



func PreOrder(t *Tree ){
	if t != nil {
		fmt.Printf("%v ",t.Data)
		PreOrder(t.Left)
		PreOrder(t.Right)

	}

}

func InOrder(t *Tree ){
	if t != nil {
		InOrder(t.Left)
		fmt.Printf("%v ",t.Data)
		InOrder(t.Right)
	}

}

func PostOrder(t *Tree ){
	if t != nil {
		PostOrder(t.Left)
		PostOrder(t.Right)
		fmt.Printf("%v ",t.Data)
	}

}


func MaxDepthWithDNF( root  *Tree) int{
	if root == nil {
		return 0
	}else{
		return max( MaxDepthWithDNF(root.Left), MaxDepthWithDNF(root.Right) )+1
	}

}



func NewNode( left *Tree, right *Tree ) *Tree {

	return  &Tree{nil,left,right}
}



func BSTLookUp( root *Tree ,data int) *Tree{
	current := root
	for  current!= nil {

		if current.Data.(int) < data {
			current = current.Right

		}else if current.Data.(int) > data {
			current = current.Left

		}else{
			return current
		}

	}
	fmt.Println("Sorry,can not find the number you input in this BST！")
	return nil



}
func main() {
	

	var it Initer
    root := NewNode(nil, nil)
	it = root
	it.SetData(20)

	//创建左子树
	a := NewNode(nil, nil)
	a.SetData(10)

	//创建右子树
	b := NewNode(nil, nil)
	b.SetData(30)

	//
	c := NewNode(nil, nil)
	c.SetData(25)

	d := NewNode(nil, nil)
	d.SetData(35)
	root.Left = a
	root.Right = b
	b.Left = c
	b.Right =  d



	fmt.Println( BSTLookUp(root,1))

}




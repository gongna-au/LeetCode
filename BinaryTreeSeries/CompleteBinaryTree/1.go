/* 
一棵高度平衡二叉树定义为：一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。
示例 1:
给定二叉树 [3,9,20,null,null,15,7]

		3
		/\
	   9  20  
      	 /  \
 		15   7


返回 true 。 
*/
/* 
二叉树中除去最后一层节点为满二叉树，且最后一层的结点依次从左到右分布，则此二叉树被称为
完全二叉树。

完全二叉树的节点个数
给出一个完全二叉树，求出该树的节点个数。
说明：
完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大
值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个
节点。
输入:


	 1
	/ \
   2   3
  / \ /
 4  5 6


输出: 6

         1
		 /\
          0
		  /\ 
		  0 1



		  1
		 /\
          0
		  /\ 
		    1


			 1
			/ \
			1  0
		   / \
		   0  0
		  / \
         0   0

*/
package main




import (
	"fmt";
	//"math"
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




func NewNode( left *Tree, right *Tree ) *Tree {

	return  &Tree{nil,left,right}
}







func CompleteBinaryTreeNum(root *Tree ) int {
	
	if root != nil {
		
		return CompleteBinaryTreeNum(root.Left)+CompleteBinaryTreeNum(root.Right)+1
		

	}
	return 0
	
}
func Abs(a int )int {
	if a >= 0{
		return a

	}else {
		return -a
	}
}
func Depth(root *Tree) int   {

	if root == nil {
		return  0
	}
	
	return Findmax(Depth(root.Left), Depth(root.Right))+1
}

func Findmax( a int ,b int ) int{

	if a >= b{
		return a

	}else{
		return b
	}

}


/* 
             1
			/ \
			1  0
		   / \
		   0  0
		  / 
         0   



		  */


func main(){
	var it Initer
    root := NewNode(nil, nil)
	it = root
	it.SetData(1)

	//创建左子树
	a := NewNode(nil, nil)
	a.SetData(1)

	//创建右子树
	b := NewNode(nil, nil)
	b.SetData(0)

	//
	c := NewNode(nil, nil)
	c.SetData(0)

	d := NewNode(nil, nil)
	d.SetData(0)

	  e := NewNode(nil, nil)
	e.SetData(0)
 
	/*f := NewNode(nil, nil)
	f.SetData(0) */

	root.Left = a
	root.Right = b
	a.Left = c
	a.Right =  d
	c.Left = e
	//c.Right = f 


	fmt.Println(CompleteBinaryTreeNum(root) )

}
	






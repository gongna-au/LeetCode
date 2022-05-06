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
假设有一棵树，最上层的是root节点，而父节点会依赖子节点。如果现在有一些节点已经标记为无效，
我们要删除这些无效节点。如果无效节点的依赖的节点还有效，那么不应该删除，如果无效节点和它的
子节点都无效，则可以删除。剪掉这些节点的过程，称为剪枝，目的是用来处理二叉树模型中的依赖问
题。
我们还是通过一道题目来进行具体学习。
题目分析
叉树的剪枝
给定二叉树根结点 root ，此外树的每个结点的值要么是 0，要么是 1。返回移除了所有不包含 1
的子树的原二叉树。( 节点 X 的子树为 X 本身，以及所有 X 的后代。)
示例1:
1 输入: [1,null,0,0,1]
2 输出: [1,null,0,null,1] 
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
	"math"
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







func BalancedBinaryTree(root *Tree) bool{
	if root == nil  {
		return true
	}
	if BalancedBinaryTree(root.Left)==false || BalancedBinaryTree(root.Right)==false {
		return false
	}
	if math.Abs( Depth(root.Left) - Depth(root.Right) )  >float64(1) {
		return false
	}
	return true
	
}
func Depth(root *Tree) float64  {

	if root == nil {
		return  float64(0)
	}
	
	return Findmax(Depth(root.Left), Depth(root.Right))+float64(1)
}

func Findmax( a float64 ,b float64 ) float64{

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

	 /* e := NewNode(nil, nil)
	e.SetData(0)
 */
	/*f := NewNode(nil, nil)
	f.SetData(0) */

	root.Left = a
	root.Right = b
	a.Left = c
	a.Right =  d
	//c.Left = e
	//c.Right = f 


	fmt.Println( BalancedBinaryTree(root) )

}
	







/*root = [5,3,6,2,4,null,7]
key = 3

    7
   / \
  6   
 /     \
2       
/\
  4
 /  \
 1	5
   /\
  1  6

    5
   / \
  3   6
 / \   \
2   4   7
    /\
  3.5 4.5


给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。
    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。
    5
   / \
  2   6
   \   \
    4   7 
*/
package main 
import (
	"fmt";
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
	fmt.Println("Sorry,can not find the number you input in this BST!")
	return nil



}
// 递归结束返回要删除的节点
func deleteNode(root *Tree, key int) *Tree{
	if root == nil {
		return nil
	}
	if key < root.Data.(int) {
		root.Left=deleteNode( root.Left, key )
		return root
	}
	if key > root.Data.(int) {
		root.Right=deleteNode( root.Right, key )
		return root
	}

	// 在递归时找到了这个节点 
	if root.Right == nil {

		return root.Left

	}

	if root.Left == nil {
		//右节点替换
		return root.Right

	}
	


		minNode := root.Right
		
		for minNode.Left != nil{
			minNode = minNode.Left
		}
		root.Data = minNode.Data
		//右子树变为删除最小节点的子树木
		root.Right = deleteMinNode(root.Right)
		
		return root

}

	


func deleteMinNode( root *Tree ) *Tree {
	
	if root.Left == nil {
		ptr:=root.Right
		root.Right = nil
		return ptr
		
	}
		//fmt.Println(PreNode)

		
			
	 root.Left = deleteMinNode((root.Left))
	//fmt.Println(basicRoot)

	return root
	 		
 
}
func main(){
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
		PreOrder(deleteNode(root,30))

		//PreOrder(deleteNode(root,30))
		/* fmt.Println(SearchMinLast(root.Right))
		

		fmt.Println(SearchKeyNode(root,30))
		PreOrder(deleteNode(root,30)) */
		/* 
		20
	   /   \
	 10    35
	       / 
		  25  


		  25
		 /   \
		10   35
		     /\
			25 35

		
		*/
}

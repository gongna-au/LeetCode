/* 
二叉树的最大深度
给定一个二叉树，找出其最大深度。二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
说明: 叶子节点是指没有子节点的节点。

示例：

给定二叉树 [3,9,20,null,null,15,7]，
    3   
   / \  
  9  20    
    /  \  
   15   7 
*/
/*    =max(4,20)+1
	  =max(max(Null,Null)+1,max(15,7)+1) +1
	  =max(max(Null,NUll)+1,max(max(Null,Null)+1,max(Null,Null)+1)+1)+1
	  =max(1,max(1,1)+1)+1
	  =max(1,1+1)+1
	  =2+1
	  =3 
*/
package main
import (
	"fmt";
	//"strconv"
	//"btree"
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



func  CreateTree( node *Tree , array []string ) *Tree {
	

	/* if len(array) == 0{
		fmt.Println("Please input an unempty array")
		
	}else {
		for {
			if i==0  {
				node = NewNode( nil, nil )
				root = node 
				v,_:=strconv.Atoi( array[i] )
				node.SetData(v)
				leftIndex := 2*i+1
				rightIndex := 2*i+2
				if leftIndex > len(array)-1 ||  rightIndex > len(array)-1 {
					break
				}
				if array[leftIndex] != ""{
					left := NewNode(nil,nil) 
					v,_ =strconv.Atoi( array[leftIndex ] )
					left.SetData( v )
					node.Left = left
				}
				if array[rightIndex] != ""{
					right := NewNode(nil,nil)
					v,_ =strconv.Atoi( array[rightIndex] )
					right.SetData( v )
					node.Right =right
				}

			}else{

			}

			if array[i] != ""{
				node = NewNode( nil, nil )
				v,_:=strconv.Atoi( array[i] )
				node.SetData(v)
				leftIndex := 2*i+1
				rightIndex := 2*i+2
			}

			
			
			if array[leftIndex] != ""{
				left := NewNode(nil,nil) 
				v,_ =strconv.Atoi( array[leftIndex ] )
				left.SetData( v )
				node.Left = left
			}
			if array[rightIndex] != ""{
				right := NewNode(nil,nil)
				v,_ =strconv.Atoi( array[rightIndex] )
				right.SetData( v )
				node.Right =right
			}

			i++
		}


	} */
	return node
	
}
func main() {
	

	//array := []string {"3", "4", "20", "", "", "15", "7"} 
	//创建根节点
	var it Initer
    root := NewNode(nil, nil)
	it = root
	it.SetData(3)

	//创建左子树
	a := NewNode(nil, nil)
	a.SetData(4)

	//创建右子树
	b := NewNode(nil, nil)
	b.SetData(20)

	//
	c := NewNode(nil, nil)
	c.SetData(15)

	d := NewNode(nil, nil)
	d.SetData(17)
	root.Left = a
	root.Right = b
	b.Left = c
	b.Right =  d

	//PreOrder(root)
	fmt.Println( MaxDepthWithDNF(root) )

}

	


	

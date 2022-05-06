package main
import (
	"fmt";
	"container/list"
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



func LeftMax(root *Tree) int {

	if (root.Right == nil){
		return root.Data.(int)
	}else{
		return LeftMax(root.Right)
	}
	
}
func RightMin(root *Tree) int {

	if (root.Right == nil){
		return root.Data.(int)
	}else{
		return RightMin(root.Left)
	}

}
// 递归求出总的节点个数
func AllTreeNodeNum( root *Tree ) int {
	if root == nil {
		return 0
	}else {
		return AllTreeNodeNum(root.Left)+ AllTreeNodeNum(root.Right)+1
	}

}

//递归求合理的节点个数
func ReasonableNodeNum( root *Tree) int {

	//层次遍历节点并判断
	 queue:= list.New()
	 queue.PushFront(root)
	 tag := 0

	 for ( queue.Len() != 0 ) {
		node := queue.Front()
		nodePtr := node.Value.(*Tree) 
		if  (nodePtr.Left == nil) && (nodePtr.Right == nil) {
			tag++
		}else if (nodePtr.Left != nil) && (nodePtr.Right == nil) {
			queue.PushBack(nodePtr.Left)
			if  (nodePtr.Left.Data.(int) < nodePtr.Data.(int)) {
				tag++
			}
			
		}else if (nodePtr.Left == nil) && (nodePtr.Right != nil) {
			queue.PushBack(nodePtr.Right)
			if (nodePtr.Right.Data.(int) > nodePtr.Data.(int)){
				tag++
			}
			
		}else if (nodePtr.Left != nil) && (nodePtr.Right != nil)  {
			if  (nodePtr.Left.Data.(int) < nodePtr.Data.(int))   && (nodePtr.Right.Data.(int) > nodePtr.Data.(int)){
				tag++
			}
			queue.PushBack(nodePtr.Left)
			queue.PushBack(nodePtr.Right)
			
		}
		queue.Remove(node)
	 }
	/* queue := list.New()
	queue.PushFront(node)
	result := []int{}
	for (queue.Len() !=0) {
		visitNode := queue.Front()
		if visitNode.Value.(*Tree).Left != nil{
			queue.PushBack( visitNode.Value.(*Tree).Left)
		}
		if visitNode.Value.(*Tree).Right != nil{
			queue.PushBack( visitNode.Value.(*Tree).Right)
		}

		// 把interface、转化为int类型
		v, _ := visitNode.Value.(*Tree).Data.(int)
		result = append(result, v)
        queue.Remove(visitNode) // 将第一个元素删
	
	} */
		return tag
}

func BSTAndVerification( root *Tree) bool {
	if root == nil {
		return false
	}else{

		if AllTreeNodeNum(root) != ReasonableNodeNum(root) {
			return false
		}else{
			//在每个节点都合理的情况下 求出左边节点的最大值 和 求出右边节点的最大值
			leftMax := LeftMax(root.Left)
			rightMin := RightMin(root.Right)	
			if ( leftMax < root.Data.(int) && rightMin > root.Data.(int) ){
				return true
			}
		}

	}
	
	return false
}




func main() {
	

	//array := []string {"5", "4", "7", "", "", "6", "8"} 
	//创建根节点
	var it Initer
    root := NewNode(nil, nil)
	it = root
	it.SetData(5)

	//创建左子树
	a := NewNode(nil, nil)
	a.SetData(4)

	//创建右子树
	b := NewNode(nil, nil)
	b.SetData(7)

	//
	c := NewNode(nil, nil)
	c.SetData(6)

	d := NewNode(nil, nil)
	d.SetData(8)
	root.Left = a
	root.Right = b
	b.Left = c
	b.Right =  d

	//PreOrder(root)
	fmt.Println( AllTreeNodeNum(root) )
	fmt.Println( ReasonableNodeNum(root) )
	fmt.Println( LeftMax(root.Left) )
	fmt.Println( RightMin(root.Right) )
	fmt.Println(BSTAndVerification(root))

}



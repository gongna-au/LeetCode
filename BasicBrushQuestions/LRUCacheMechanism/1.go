package main

type LinkNode struct {
	key  interface{}
	val  interface{}
	pre  *LinkNode
	next *LinkNode
}
type LRUCache struct {
	mapList map[int]*LinkNode
	cap     int
	head    *LinkNode
	tail    *LinkNode
}

func Constructor(capacity int) *LRUCache {
	head := &LinkNode{
		nil,
		nil,
		nil,
		nil,
	}
	tail := &LinkNode{
		nil,
		nil,
		nil,
		nil,
	}

	head.next = tail
	tail.pre = head
	mapList := map[int]*LinkNode{}

	return &LRUCache{
		mapList,
		capacity,
		head,
		tail,
	}

}
func (this *LRUCache) Get(key int) int {
	cache := this.mapList
	if k, exist := cache[key]; exist {

		this.RemoveNode(k)
		this.MoveTohead(k)
		if temp, ok := k.val.(int); ok {
			return temp
		}

	} else {
		node := &LinkNode{
			key,
			nil,
			nil,
			nil,
		}
		this.AdjustCapacity()
		this.MoveTohead(node)

	}

	return -1

}

func (this *LRUCache) RemoveNode(node *LinkNode) {

	node.pre.next = node.next
	node.next.pre = node.pre

}

func (this *LRUCache) MoveTohead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)

}

func (this *LRUCache) AddNode(node *LinkNode) {

	head := this.head

	node.next = head.next
	head.next.pre = node

	head.next = node
	node.pre = head

}

func (this *LRUCache) Put(key int, value int) {

	cache := this.mapList

	if k, ok := cache[key]; ok {
		k.val = value
		this.RemoveNode(k)
		this.MoveTohead(k)
		this.AdjustCapacity()
	} else {
		node := &LinkNode{
			key,
			value,
			nil,
			nil,
		}

		this.MoveTohead(k)
		cache[key] = node
		this.AdjustCapacity()
	}

}
func (this *LRUCache) AdjustCapacity() {
	if len(this.mapList) == this.cap {
		this.RemoveNode(this.tail)
	}

}

func main() {

}

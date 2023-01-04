package stack

import "fmt"

/*
   基于链表实现栈
   时间复杂度都是O(1)
   都是基于特定位置的操作
 */

type node struct {
	next *node
	val  int
}

type LinkedListStack struct {
	//栈顶结点
	topNode *node
}

func NewLinkedListStack() *LinkedListStack{
	return &LinkedListStack{nil}
}

func (this *LinkedListStack) IsEmpty() bool{
	return this.topNode == nil
}

func (this *LinkedListStack) Push(v int){
	this.topNode = &node{next: this.topNode,val: v}
}

func (this *LinkedListStack) Pop() int {
	if this.IsEmpty(){
		return 0
	}
	v := this.topNode.val
	this.topNode = this.topNode.next
	return v
}

func (this *LinkedListStack) Top() int{
	if this.IsEmpty(){
		return 0
	}
	return this.topNode.val
}

func (this *LinkedListStack) Flush(){
	this.topNode = nil
}

func (this *LinkedListStack) Print() {
	if this.IsEmpty(){
		fmt.Println("empty stack")
	}else {
		cur := this.topNode
		for nil != cur {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
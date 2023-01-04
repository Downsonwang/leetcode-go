package ArrayQueue

import "fmt"

// 基于链表的队列实现
type ListNode struct {
	val interface{}
	next *ListNode
}

type LinkedLIstQueue struct {
	head *ListNode
	tail *ListNode
	length int
}

func NewLinkedListQueue() *LinkedLIstQueue{
	return &LinkedLIstQueue{nil,nil,0}
}

// 入队操作
func (this *LinkedLIstQueue) EnQueue(v interface{}){
	node := &ListNode{v,nil}
	if nil == this.tail {
		this.tail = node
		this.head = node
	}else {
		this.tail.next = node
		this.tail = node
	}
	this.length++
}

func (this *LinkedLIstQueue) DeQueue() interface{}  {
	if this.head == nil {
		return nil
	}
	v := this.head.val
	this.head= this.head.next
	this.length--
	return v
}

func (this *LinkedLIstQueue) String() string{
	if this.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := this.head; cur != nil ; cur = cur.next{
		result += fmt.Sprintf("<-%+v",cur.val)
	}
	result += "<-tail"
	return result
}
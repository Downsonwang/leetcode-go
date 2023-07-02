package _41designcirculardeque

type node struct {
	prev,next *node
	val int
}
type CircularDeque struct {
	head,tail *node
	cap,size int
}

func MConstructor(k int) CircularDeque {
	return CircularDeque{cap: k}
}


func (this *CircularDeque) InsertFront(value int) bool  {
	if this.IsFull() {
		return false
	}
	node := &node{val: value}
	if this.IsEmpty(){
		this.head = node
		this.tail = node
	}else{
		node.next = this.head
		this.head.prev = node
	    this.head = node
	}
	this.size++
	return true

}

func (this *CircularDeque) InsertLast(value int) bool  {
	if this.IsFull(){
		return false
	}
	node := &node{val: value}
	if this.IsEmpty(){
		this.head = node
		this.tail = node
	}else{
		this.tail.next = node
		node.prev = this.tail
		this.tail = node
	}
	this.size++
	return true
}

func (this *CircularDeque) DeleteFront() bool {
	if this.IsEmpty(){
		return false
	}

	//修改队头指针,将其加一,表示删除队头元素 因为是
	this.head = this.head.next
	if this.head != nil {
		this.head.prev = nil
	}
	this.size--
	return true
}


func (this *CircularDeque) DeleteLast() bool {
	if this.IsEmpty(){
		return false
	}
	this.tail = this.tail.prev
	if this.tail != nil{
		this.tail.next = nil
	}
	this.size--
	return true
}


func (this *CircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.val
}


func (this *CircularDeque) GetRear() int {
	if this.IsEmpty(){
		return -1
	}
	return this.tail.val
}


func (this *CircularDeque) IsEmpty() bool {
	return this.size == 0
}


func (this *CircularDeque) IsFull() bool {
   return this.size == this.cap
}
